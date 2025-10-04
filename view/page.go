package view

import (
	"fmt"
	"io"
	"strings"
)

var header = `<html>
	<head>
		<style>
			@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap');
		</style>
		<link rel="stylesheet" href="static/main.css">
	</head>	
<body>
<main>
<h1>dsadsa</h1>
`

func RenderPage(pageComponent Component, w io.Writer) error {
	sb := strings.Builder{}

	_, err := sb.WriteString(header)
	if err != nil {
		return fmt.Errorf("writing page header: %w", err)
	}

	if pageComponent != nil {
		err = pageComponent.Render(&sb)
		if err != nil {
			return fmt.Errorf("writing page body: %w", err)
		}
	}

	_, err = sb.WriteString(`</main></body>`)
	if err != nil {
		return fmt.Errorf("writing page end: %w", err)
	}

	_, err = io.WriteString(w, sb.String())
	if err != nil {
		return fmt.Errorf("writing page to writer: %w", err)
	}

	return nil
}
