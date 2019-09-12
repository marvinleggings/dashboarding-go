package main

import (
	"log"
	"encoding/json"
)

var b = []byte(`{"Layouts":[                                                                       
                   {"x": "0", "y" : "0", "w" : "5", "h" : "5", "content" : "html" },
                   {"x": "1", "y" : "1", "w" : "10", "h" : "10", "content" : "html1" },
                   {"x": "0", "y" : "0", "w" : "5", "h" : "5", "content" : "html" }
               ]}`)

func main() {
	var f interface{}
	   	err := json.Unmarshal(b, &f)
		if err != nil {
	        panic("OMG!")
	    }
	m := f.(map[string]interface{})
	
	for key, value := range m{
		log.Println(key, value)
		// log.Println(value.([]interface{})[1].(map[string]interface{})["x"])
		switch val := value.(type) {
		
		case []interface{}:
			for n := range val {
				log.Println(val[0], n)
			}
		}

	}
}

//value["persons"][0]["name"]