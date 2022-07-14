{{ define "myalert" }}
  [{{.Status}}] {{ .Labels.alertname }}
  Labels:
  {{ range .Labels.SortedPairs }}
    {{ .Name }}: {{ .Value }}
  {{ end }}
{{ end }}