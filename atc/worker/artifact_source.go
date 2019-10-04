package worker

import (
	"context"
	"io"

	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter . ArtifactSource

// Source represents data produced by the steps, that can be transferred to
// other steps.
type ArtifactSource interface {
	// StreamTo copies the data from the source to the destination. Note that
	// this potentially uses a lot of network transfer, for larger artifacts, as
	// the ATC will effectively act as a middleman.
	StreamTo(context.Context, lager.Logger, ArtifactDestination) error

	// StreamFile returns the contents of a single file in the artifact source.
	// This is used for loading a task's configuration at runtime.
	//
	// If the file cannot be found, FileNotFoundError should be returned.
	StreamFile(context.Context, lager.Logger, string) (io.ReadCloser, error)

	// VolumeOn attempts to locate a volume equivalent to this source on the
	// given worker. If a volume can be found, it will be used directly. If not,
	// `StreamTo` will be used to copy the data to the destination instead.
	VolumeOn(lager.Logger, Worker) (Volume, bool, error)
}