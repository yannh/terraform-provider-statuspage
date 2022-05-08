
resource "statuspage_component" "my_component" {
  page_id     = "j1r06kbq73f0"
  name        = "Component1"
  description = "Status of component1"
  status      = "operational"
}


resource "statuspage_component_group" "my_group" {
  page_id     = "j1r06kbq73f0"
  name        = "terraform"
  description = "Created by terraform"
  components  = ["${statuspage_component.my_component.id}"]
}
