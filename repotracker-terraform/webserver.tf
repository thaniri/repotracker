provider "aws" {
  shared_credentials_file = "${var.shared_credentials_file}"
  profile = "${var.profile}"
  region = "${var.region}"
}

resource "aws_security_group" "webserver_inbound" {
  name = "${var.webserver_inbound_SG_name}"
  description = "allow HTTP/S and SSH from the internet"

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress { 
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Direction = "inbound"
    Role = "webserver"
  }
}

resource "aws_instance" "webserver001" {
  ami = "ami-0bbe6b35405ecebdb"
  instance_type = "t3.micro"
  key_name = "terraform"
  security_groups = [
    "${var.webserver_inbound_SG_name}"
  ]
}

