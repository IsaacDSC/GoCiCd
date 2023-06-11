variable "image" {
  type    = string
  default = "latest"
  description = "The docker image used for task."
  validation {
    condition     = substr(var.image, 0, 27) == "rodrigoresende/oli-bff-sdk:"
    error_message = "Not Permission Deployment The Image."
  }
}
  
job "oli-bff-1" {
  type = "service"


  group "bff-sdk-group" {
    count = 1

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

    network {

      port "application" {
        to = 3000
        // static = 8093
        // static = 8092
      }

      // port "ingress" {
      //   to = 80
      //   static = 8080
      // }
    }

    restart {
    attempts = 10
    delay    = "30s"
    mode     = "fail"
    }

    service {
      name     = "oli-bff-1"
      port     = "application"
      provider = "nomad"

      check_restart {
        limit           = 3
        grace           = "10s"
        ignore_warnings = false
      }

      check {
        name     = "healtcheck"
        type     = "http"
        path     = "/api/v1/status"
        interval = "5s"
        timeout  = "1s"
        method   = "GET"
        header {
          Authorization = ["Basic ZWxhc3RpYzpjaGFuZ2VtZQ=="]
        }
      }
    }

    // update {
    //   max_parallel     = 1
    //   canary           = 1
    //   min_healthy_time = "30s"
    //   auto_revert      = true
    //   auto_promote     = false
    // }

    task "app-bff-sdk-task" {
      driver = "docker"

      config {
        // image = "rodrigoresende/oli-bff-sdk:${image}"
        image = var.image
        ports = ["application"]
      }
      resources {
        cpu    = 500 # MHz
        memory = 500 # MB
      }
    }

    // task "ingress-bff-sdk-task" {
    //   driver = "docker"

    //   config {
    //     image = "nginx:latest"
    //     ports = ["ingress"]
    //     volumes = [
    //         "local:/etc/nginx/conf.d",
    //     ]
    //   }

    //   resources {
    //     cpu    = 50 # MHz
    //     memory = 50 # MB
    //   }

    //   template {
    //     data = <<EOF
    //     server {
    //       listen 80;

    //       location / {
    //           proxy_pass https://api.oli.services/bff;
    //       }
    //     }
    //     EOF

    //     destination   = "local/load-balancer.conf"
    //     change_mode   = "signal"
    //     change_signal = "SIGHUP"
    //   }
    // }

  }
}