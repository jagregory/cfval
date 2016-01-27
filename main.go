package main

import "os"
import "io/ioutil"
import "encoding/json"
import "fmt"

type Resource interface {
	Validate(t Template, context []string) (bool, []Failure)
}

func main() {
  bytes, err := ioutil.ReadAll(os.Stdin)

  if err != nil {
    fmt.Println("Error reading JSON")
    return
  }

	var temp struct {
		Resources map[string]struct {
			Type       string
			Properties json.RawMessage
		}
	}

	err = json.Unmarshal(bytes, &temp)

	if err != nil {
		fmt.Println(err)
		return
	}

	template := &Template{
		Resources: make(map[string]Resource),
	}

	for key, rawResource := range temp.Resources {
		if rawResource.Type == "AWS::EC2::Subnet" {
			var res Aws_Ec2_Subnet
			err = json.Unmarshal(rawResource.Properties, &res)
			if err != nil {
				fmt.Println(err)
				return
			}
			template.Resources[key] = res
		}
	}

	if ok,errors := template.Validate(); !ok {
		for _,err := range errors {
			fmt.Println(err)
		}
	}
}
