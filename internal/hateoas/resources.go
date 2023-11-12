package hateoas

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jtonynet/cine-catalogo/internal/utils"
	"github.com/pmoule/go2hal/hal"
	"github.com/pmoule/go2hal/halforms"
)

// WRAPPER FOR go2hal/hal AND go2hal/halforms TO SIMPLIFY USE
// https://rwcbook.github.io/hal-forms/#_the_hal_forms_media_type
// https://github.com/pmoule/go2hall
// https://hal-explorer.com/#theme=Dark&allHttpMethodsForLinks=true&hkey0=Accept&hval0=application/prs.hal-forms+json&uri=http://localhost:8080/v1/

type resource struct {
	name         string
	resourceURL  string
	linkRelation hal.LinkRelation
	template     halforms.Template
}

func NewResource(name, segmentURL, httpMethod string) (*resource, error) {
	resourceURL := fmt.Sprintf("%s/%s", rootURL, segmentURL)
	linkRelation, err := halforms.NewHALFormsRelation(name, resourceURL)
	if err != nil {
		return nil, err
	}

	r := &resource{
		name:         name,
		resourceURL:  resourceURL,
		linkRelation: linkRelation,
	}
	r.template = r.newTemplate(httpMethod)

	return r, nil
}

func (r *resource) newTemplate(httpMethod string) halforms.Template {
	template := halforms.NewTemplate()
	template.Method = httpMethod
	template.Target = r.resourceURL
	template.Key = r.name
	template.Title = ""

	return *template
}

func (r *resource) RequestToProperties(request interface{}) error {
	hateoasTag := "hateoas"

	val := reflect.ValueOf(request)
	if val.Kind() != reflect.Struct {
		return errors.New("input should be a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get(hateoasTag)

		if tag != "" {
			firstLetterLower := strings.ToLower(field.Name[:1])
			propertyName := firstLetterLower + field.Name[1:]

			property := halforms.NewProperty(propertyName)
			property.Prompt = field.Name

			propertiesStructTag := strings.Split(tag, ";")
			for _, propertyStr := range propertiesStructTag {
				propKeyValue := strings.SplitN(propertyStr, ":", 2)

				switch propKeyValue[0] {
				case "name":
					property.Name = propKeyValue[1]
				case "prompt":
					property.Prompt = propKeyValue[1]
				case "placeholder":
					property.Placeholder = propKeyValue[1]
				case "required":
					required, _ := utils.StringParseBoolean(propKeyValue[1])
					property.Required = required
				}

			}
			r.template.Properties = append(r.template.Properties, property)
		}
	}

	return nil
}
