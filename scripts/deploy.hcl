variable "image" {
  type    = string
  default = "latest"
  description = "The docker image used for task."
//   validation {
//     condition     = substr(var.image, 0, 27) == "rodrigoresende/oli-bff-sdk:"
//     error_message = "Not Permission Deployment The Image."
//   }
}

job "test-webapp-go" {
  datacenters = ["dc1"]


  group "webapp-group" {
    count = 3
    network {
      port "http" {
        to = 3000
      }
    }


    restart {
      attempts = 10
      delay    = "30s"
      mode     = "fail"
    }


    scaling {
      enabled = true
      min     = 1
      max     = 10


      policy {
        check "96pct" {
          strategy "app-sizing-percentile" {
            percentile = "70"
          }
        }
      }
    }


    service {
      name = "test-webapp-go"
      port = "http"


      check {
        type     = "http"
        path     = "/"
        interval = "2s"
        timeout  = "2s"
        header {
          Authorization = ["Basic ZWxhc3RpYzpjaGFuZ2VtZQ=="]
        }
      }
    }


    task "webapp-task" {
      env {
        PORT    = "${NOMAD_PORT_http}"
        NODE_IP = "${NOMAD_IP_http}"
      }


      driver = "docker"


      config {
        // image = "isaacdsc/test-webapp-go:latest"
        image = var.image
        ports = ["http"]
      }


      resources {
        cpu    = 2000 # MHz
        memory = 500 # MB
      }
    }
  }
}