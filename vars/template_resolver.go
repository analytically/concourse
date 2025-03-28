package vars

// TemplateResolver processes templates with variable interpolation.
type TemplateResolver struct {
	template []byte
	params   []Variables
}

// NewTemplateResolver creates a template resolver with a template and variable sources.
// When multiple param sources are provided, they are tried in order for variable lookup.
// The first source that contains a variable will be used for that variable.
// See implementation of NewMultiVars for details.
func NewTemplateResolver(configPayload []byte, params []Variables) TemplateResolver {
	return TemplateResolver{
		template: configPayload,
		params:   params,
	}
}

// Resolve processes the template by replacing variable placeholders with their values.
// If expectAllKeys is true, an error is returned when variables cannot be resolved.
func (resolver TemplateResolver) Resolve(expectAllKeys bool) ([]byte, error) {
	var err error

	resolver.template, err = resolver.resolve(expectAllKeys)
	if err != nil {
		return nil, err
	}

	return resolver.template, nil
}

// resolve handles the actual template resolution using the provided variable sources.
func (resolver TemplateResolver) resolve(expectAllKeys bool) ([]byte, error) {
	tpl := NewTemplate(resolver.template)
	bytes, err := tpl.Evaluate(NewMultiVars(resolver.params), EvaluateOpts{ExpectAllKeys: expectAllKeys})
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
