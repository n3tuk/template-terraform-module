package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
)

func TestModuleExamples(t *testing.T) {
	t.Parallel()

	tests := {% raw %}[]string{{% endraw %}{% for name in cookiecutter.components.examples %}"{{ name }}"{% if not loop.last %}, {% endif %}{% endfor %}{% raw %}}{% endraw %}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			id := tc
			tag := random.UniqueId()
			dir := "../examples/" + id

			e := NewExampleConfiguration(t, dir, tag)
			defer e.Destroy(t)

			e.Init(t)
			e.Apply(t)
		})
	}
}
