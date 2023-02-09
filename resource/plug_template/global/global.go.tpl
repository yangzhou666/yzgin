package global

{{- if .HasGlobal }}

import "yzgin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}