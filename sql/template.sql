SELECT 
	id,
	name,
	color,
	price,
FROM table 
WHERE 
    1=1
    {{/* Name */}}
    {{- if not .Name.Skip -}}
    AND name = '{{ .Name.Value}}'
    {{- end -}}
    {{/* Color */}}
    AND color IN (
            {{- range  $i, $v := .Color -}}
                {{- if ne $i 0}},{{end -}}'{{$v}}'
            {{- end -}}
        )
;