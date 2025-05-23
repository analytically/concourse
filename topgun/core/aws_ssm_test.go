package topgun_test

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"

	. "github.com/concourse/concourse/topgun"
	. "github.com/concourse/concourse/topgun/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS SSM", func() {
	var ssmAPI *ssm.Client
	var awsRegion string
	var awsCreds aws.Credentials

	BeforeEach(func(ctx SpecContext) {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			Skip("can not load default AWS config")
		}

		ssmAPI = ssm.NewFromConfig(cfg)
		awsRegion = cfg.Region
		awsCreds, err = cfg.Credentials.Retrieve(ctx)
		if err != nil {
			Skip("can not retrieve AWS credentials")
		}
	})

	Describe("A deployment with SSM", func(ctx SpecContext) {
		BeforeEach(func() {
			sessionToken := awsCreds.SessionToken
			if sessionToken == "" {
				sessionToken = `""`
			}
			Deploy(
				"deployments/concourse.yml",
				"-o", "operations/configure-ssm.yml",
				"-v", "aws_region="+awsRegion,
				"-v", "aws_access_key="+awsCreds.AccessKeyID,
				"-v", "aws_secret_key="+awsCreds.SecretAccessKey,
				"-v", "aws_session_token="+sessionToken,
			)
		})

		Context("/api/v1/info/creds", func() {
			type responseSkeleton struct {
				Ssm struct {
					AwsRegion string `json:"aws_region"`
					Health    struct {
						Response struct {
							Status string `json:"status"`
						} `json:"response"`
						Error string `json:"error,omitempty"`
					} `json:"health"`
					PipelineSecretTemplate string `json:"pipeline_secret_template"`
					TeamSecretTemplate     string `json:"team_secret_template"`
				} `json:"ssm"`
			}

			var (
				atcURL         string
				parsedResponse responseSkeleton
			)

			BeforeEach(func() {
				atcURL = "http://" + JobInstance("web").IP + ":8080"
			})

			JustBeforeEach(func() {
				token, err := FetchToken(atcURL, AtcUsername, AtcPassword)
				Expect(err).ToNot(HaveOccurred())

				body, err := RequestCredsInfo(atcURL, token.AccessToken)
				Expect(err).ToNot(HaveOccurred())

				err = json.Unmarshal(body, &parsedResponse)
				Expect(err).ToNot(HaveOccurred())
			})

			It("contains ssm config", func() {
				Expect(parsedResponse.Ssm.AwsRegion).To(Equal(awsRegion))
				Expect(parsedResponse.Ssm.Health).ToNot(BeNil())
				Expect(parsedResponse.Ssm.Health.Error).To(BeEmpty())
				Expect(parsedResponse.Ssm.Health.Response).ToNot(BeNil())
				Expect(parsedResponse.Ssm.Health.Response.Status).To(Equal("UP"))
			})
		})

		testCredentialManagement(func() {
			secrets := map[string]string{
				"/concourse-topgun/main/team_secret":                              "some_team_secret",
				"/concourse-topgun/main/pipeline-creds-test/assertion_script":     assertionScript,
				"/concourse-topgun/main/pipeline-creds-test/canary":               "some_canary",
				"/concourse-topgun/main/pipeline-creds-test/resource_type_secret": "some_resource_type_secret",
				"/concourse-topgun/main/pipeline-creds-test/resource_secret":      "some_resource_secret",
				"/concourse-topgun/main/pipeline-creds-test/job_secret/username":  "some_username",
				"/concourse-topgun/main/pipeline-creds-test/job_secret/password":  "some_password",
				"/concourse-topgun/main/pipeline-creds-test/resource_version":     "some_exposed_version_secret",
			}

			for name, value := range secrets {
				_, err := ssmAPI.PutParameter(ctx, &ssm.PutParameterInput{
					Name:      aws.String(name),
					Value:     aws.String(value),
					Type:      types.ParameterTypeSecureString,
					Overwrite: aws.Bool(true),
				})
				Expect(err).To(BeNil())
			}
		}, func() {
			secrets := map[string]string{
				"/concourse-topgun/main/team_secret":      "some_team_secret",
				"/concourse-topgun/main/resource_version": "some_exposed_version_secret",
			}

			for name, value := range secrets {
				_, err := ssmAPI.PutParameter(ctx, &ssm.PutParameterInput{
					Name:      aws.String(name),
					Value:     aws.String(value),
					Type:      types.ParameterTypeSecureString,
					Overwrite: aws.Bool(true),
				})
				Expect(err).To(BeNil())
			}
		})
	})
})
