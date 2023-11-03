package config

var (
	Version       = "0.2.0"
	WireCmd       = "github.com/google/wire/cmd/wire@latest"
	GoalCmd       = "github.com/xushuhui/goal@latest"
	RepoBase      = "https://github.com/xushuhui/goal-layout-base.git"
	RepoAdvanced  = "https://github.com/xushuhui/goal-layout-advanced.git"
	RunExcludeDir = ".git,.idea,tmp,vendor"
	RunIncludeExt = "go,html,yaml,yml,toml,ini,json,xml,tpl,tmpl"
)
