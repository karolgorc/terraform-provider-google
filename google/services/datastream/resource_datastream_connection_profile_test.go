// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package datastream_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDatastreamConnectionProfile_update(t *testing.T) {
	// this test uses the random provider
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	random_pass_1 := acctest.RandString(t, 10)
	random_pass_2 := acctest.RandString(t, 10)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDatastreamConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamConnectionProfile_update(context),
			},
			{
				ResourceName:            "google_datastream_connection_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location"},
			},
			{
				Config: testAccDatastreamConnectionProfile_update2(context, true),
			},
			{
				ResourceName:            "google_datastream_connection_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "postgresql_profile.0.password"},
			},
			{
				// Disable prevent_destroy
				Config: testAccDatastreamConnectionProfile_update2(context, false),
			},
			{
				Config: testAccDatastreamConnectionProfile_mySQLUpdate(context, true, random_pass_1),
			},
			{
				ResourceName:            "google_datastream_connection_profile.mysql_con_profile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "mysql_profile.0.password"},
			},
			{
				// run once more to update the password. it should update it in-place
				Config: testAccDatastreamConnectionProfile_mySQLUpdate(context, true, random_pass_2),
			},
			{
				ResourceName:            "google_datastream_connection_profile.mysql_con_profile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "mysql_profile.0.password"},
			},
			{
				// Disable prevent_destroy
				Config: testAccDatastreamConnectionProfile_mySQLUpdate(context, false, random_pass_2),
			},
		},
	})
}

func TestAccDatastreamConnectionProfile_sshKey_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	randomPubKey1 := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCjXhptfWIrtflLZ1WeOsjCfHSEKvui0fdNXTqpqIA+2NNlFjwKS4mV3bDJIRlC5FdWG/D5LW4kvSmcTx1eSLUcvqw3i3F73Ii35AR1Rid1bY0LCBYUUgkDKyvZgDzrM7g+MwBtthoud8Axt9/bh28qtzSVNvWfxIYsa2CwtqlkZr5c6Qb6N2B9kxW8WFsCnoAeBaZDMq+LVBRsRJvBBrJm/qhMNPd07Al7wGLEnNPWmwjFT7B12sMjNr7ZNLfI9VckEyUSx3AGBFH7RImeYiWb6vZA9v5DE7kBrCoHtJK5IN9dvqEWXrrDT7RTFXd55xQqT70eZiIDNz1nexDw8ZCn user`
	randomPrivKey1 := `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEAo14abX1iK7X5S2dVnjrIwnx0hCr7otH3TV06qaiAPtjTZRY8CkuJ
ld2wySEZQuRXVhvw+S1uJL0pnE8dXki1HL6sN4txe9yIt+QEdUYndW2NCwgWFFIJAysr2Y
A86zO4PjMAbbYaLnfAMbff24dvKrc0lTb1n8SGLGtgsLapZGa+XOkG+jdgfZMVvFhbAp6A
HgWmQzKvi1QUbESbwQayZv6oTDT3dOwJe8BixJzT1psIxU+wddrDIza+2TS3yPVXJBMlEs
dwBgRR+0SJnmIlm+r2QPb+QxO5AawqB7SSuSDfXb6hFl66w0+0UxV3eecUKk+9HmYiAzc9
Z3sQ8PGQpwAAA8B2IBoLdiAaCwAAAAdzc2gtcnNhAAABAQCjXhptfWIrtflLZ1WeOsjCfH
SEKvui0fdNXTqpqIA+2NNlFjwKS4mV3bDJIRlC5FdWG/D5LW4kvSmcTx1eSLUcvqw3i3F7
3Ii35AR1Rid1bY0LCBYUUgkDKyvZgDzrM7g+MwBtthoud8Axt9/bh28qtzSVNvWfxIYsa2
CwtqlkZr5c6Qb6N2B9kxW8WFsCnoAeBaZDMq+LVBRsRJvBBrJm/qhMNPd07Al7wGLEnNPW
mwjFT7B12sMjNr7ZNLfI9VckEyUSx3AGBFH7RImeYiWb6vZA9v5DE7kBrCoHtJK5IN9dvq
EWXrrDT7RTFXd55xQqT70eZiIDNz1nexDw8ZCnAAAAAwEAAQAAAQAnvU5kb+mfhGaeBwb2
tIn9dVTKicIoezbTJOiOOKTppMjXgC8euf0/7WuBoYGJmg38rlNR6dEvMqyaj0wvkTQtR9
yQrmTuoljHkrna5TPYBswWcOMeEk6K7Md/4wfulugsiS+DgJah0xN3hKj5t9o848/wtCvP
r3iL+ZrNocFW4Ju+QrArFWTLFuJL4uc69ykgWE7I5Qkm+3Lg6aSoNazMzCu9rCblduetJq
EilQ6AOkv68xTOQ1EDIQc8xr6u6GCUvVVBwYaR3cYV6fWeLWJATqUODkEXdDZfgUerf4Io
3KirdRf0YFyJiHJh4AqWd76jWCkhCwrREx0lfMCZghoxAAAAgHwOfMJtd4wOug2BPKu0SA
HSwQ+yTTibg2xuENstd8akJC3VsU5GC8pngNAyoFpSt3QDlLpvqPqXVJSkkMbUtnPO0SIR
5ffMB97kFvNkMNDUIalwxR9DV1CMPTAnTO7NSfO8UUKRjKivpmpS6ptMjxUM0hPoDBebhx
P37In1a2jDAAAAgQDVCaoMFjHRGds1JaVjm6YviR0C2OsE55GOS7cW+I3SE63DumfHsN8i
r/u5oEQUelaauYVmi9tT3L4lReFX2tYqtyE0mbPUXcY5XfmBxBsjW1sQ6YyHlN/vGLgo33
NZZFpIg2FknTzM4qeddfbyKuqAJX27f7RrSZCf+WrJUKDWqwAAAIEAxFAn6d9na7uHnb31
TQ8PoTvkH7fwugXuG7ACLCTl3PpOSGPQAPI8rCaGOMd+uU1Jyjt3TcdPYlNAtiFQCxWLMH
RNFfeqviC85H6WzQNezNj45QqKTf5gRdHVu2NMRwn2pJjRgdIvsUaL1AY4sC0AivoEMlpx
rQYvdaDG7KsYXfUAAAAEdXNlcgECAwQFBgc=
-----END OPENSSH PRIVATE KEY-----`

	randomPubKey2 := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDmc1i/FqnVtYsTzb6LmoUGom8ISnfRCPTIFf3LLIyRFgO+qD6Dnqn5p2lLE8ksdooAGJ+EyJtV5c+3kYGnjzzH4TlB2pkt562BntrggvJ98sELQbHEDiemiLnJqqIESk5FcSXdcJ/UX/AdkbXLjSR5M8+cGGqKSb0HSnKfOWkjWwZwp/JwbvyWPIJ6IQNKzAS5HVU/J+u8ezhPd1iBdezvAuPlihpjMGQg1KW3APZoELS6/BSMpXcvDy+TwuggEPPZ0Up09BJRtqesHiZur6CnqUIzJcCWCfi5C8IfHzlhawry+iA1V5Lh06Mz7OaySXpf902RITfh+KcLxcSSMmPl user`
	randomPrivKey2 := `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEA5nNYvxap1bWLE82+i5qFBqJvCEp30Qj0yBX9yyyMkRYDvqg+g56p
+adpSxPJLHaKABifhMibVeXPt5GBp488x+E5QdqZLeetgZ7a4ILyffLBC0GxxA4npoi5ya
qiBEpORXEl3XCf1F/wHZG1y40keTPPnBhqikm9B0pynzlpI1sGcKfycG78ljyCeiEDSswE
uR1VPyfrvHs4T3dYgXXs7wLj5YoaYzBkINSltwD2aBC0uvwUjKV3Lw8vk8LoIBDz2dFKdP
QSUbanrB4mbq+gp6lCMyXAlgn4uQvCHx85YWsK8vogNVeS4dOjM+zmskl6X/dNkSE34fin
C8XEkjJj5QAAA8CppfYQqaX2EAAAAAdzc2gtcnNhAAABAQDmc1i/FqnVtYsTzb6LmoUGom
8ISnfRCPTIFf3LLIyRFgO+qD6Dnqn5p2lLE8ksdooAGJ+EyJtV5c+3kYGnjzzH4TlB2pkt
562BntrggvJ98sELQbHEDiemiLnJqqIESk5FcSXdcJ/UX/AdkbXLjSR5M8+cGGqKSb0HSn
KfOWkjWwZwp/JwbvyWPIJ6IQNKzAS5HVU/J+u8ezhPd1iBdezvAuPlihpjMGQg1KW3APZo
ELS6/BSMpXcvDy+TwuggEPPZ0Up09BJRtqesHiZur6CnqUIzJcCWCfi5C8IfHzlhawry+i
A1V5Lh06Mz7OaySXpf902RITfh+KcLxcSSMmPlAAAAAwEAAQAAAQEAq2opHRpSgfBj3vsv
PNBXGrRAOr6JmSc8TIhvG22rsU/awTqMJYMjk9v+6iVxgm06ARBPt4kwYhhrBXRqKKTW5S
aWXHGpdwfZe40Z6d39Wcnz5debzuVogOs6ptMRaHeM+QJM1AYuHN6v0I7N1vbJpo3vY4CV
3v8yZ/XshJtDpVNqHFuCh1r07aW4NlqoTy5TEvWD1VPCqAVwTLWuNMfWRGYbwqJrRUxuu3
6vqddE8yMONYMwVRKPADj0DTi3i+LK3v6QfJlxb09EhqJPOOXM+fBVzUWkUXlPjvMP4uUH
/zRrGscSI93n0V/H3/XTOJTskdEZUEFpeFbUXIphloCKEQAAAIEA9CJapVXG9HcKimXX3I
OQdwPoKONM52KnAoWjGO1N5ECydjz2yHQkNJNLFwAUefmKVy0/ce0EdyEJjoHKvCwoTWL6
3CPlWQY+7pk0Fr62iT7UjjGwCtmHB6B5G4qUlsBkVN3WCwfmBwYrziRR+qcS8hSS7m37Uy
rMbGGIHHVGPzIAAACBAP6ouUUlIN7jLdLxyApj1Cx7oW7Gp33j3goXn82WVv6+ubPJymVD
u7zmoWWVegOngoPlR1q/mHBGoB1Ec1Im5IaN5qzVrxVKraJz5Q1XRc/azpkYb1FaDFBW2O
iDaP5PHvNQpYcmE82Dg8bUqa7tYIUgq2vqHJdBZC5IvnYnGrWbAAAAgQDnqf2DVITbK5jK
UJqEmni0YE8PD3PuPGRWLmZeOcxshHR1nQIeUoXWAhCS9G7Rl5Kdr1IXzSln22OvUXMPmE
gZLd7QJVyRQ0bXhYf8nIs/UGhjq83OSoS4iSwHeZ1CrKWmVP74/+Na6fDdfJ65Z8+I4ktM
QC3v6moZVb2wrgGkfwAAAAR1c2VyAQIDBAU=
-----END OPENSSH PRIVATE KEY-----`

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "registry.terraform.io/hashicorp/time",
			},
		},
		CheckDestroy: testAccCheckDatastreamConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamConnectionProfile_sshKey_update(context, true, randomPrivKey1, randomPubKey1),
			},
			{
				ResourceName:            "google_datastream_connection_profile.ssh_connectivity_profile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "create_without_validation", "forward_ssh_connectivity.0.private_key", "postgresql_profile.0.password"},
			},
			{
				PreConfig: func() {
					fmt.Println("Waiting before proceeding to the next step...")
					time.Sleep(150 * time.Second) // Delay before the next step
				},
				Config: testAccDatastreamConnectionProfile_sshKey_update(context, true, randomPrivKey2, randomPubKey2),
			},
			{
				ResourceName:            "google_datastream_connection_profile.ssh_connectivity_profile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "create_without_validation", "forward_ssh_connectivity.0.private_key", "postgresql_profile.0.password"},
			},
			{
				PreConfig: func() {
					fmt.Println("Waiting before proceeding to the next step...")
					time.Sleep(150 * time.Second) // Delay before the next step
				},
				Config: testAccDatastreamConnectionProfile_sshKey_update(context, false, randomPrivKey2, randomPubKey2),
			},
		},
	})
}

func testAccDatastreamConnectionProfile_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-my-profile%{random_suffix}"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}
	lifecycle {
		prevent_destroy = true
	}
}
`, context)
}

func testAccDatastreamConnectionProfile_update2(context map[string]interface{}, preventDestroy bool) string {
	context["lifecycle_block"] = ""
	if preventDestroy {
		context["lifecycle_block"] = `
		lifecycle {
			prevent_destroy = true
		}`
	}
	return acctest.Nprintf(`
resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-database-instance%{random_suffix}"
    database_version = "POSTGRES_14"
    region           = "us-central1"
    settings {
      tier = "db-f1-micro"

      ip_configuration {

        // Datastream IPs will vary by region.
        authorized_networks {
            value = "34.71.242.81"
        }

        authorized_networks {
            value = "34.72.28.29"
        }

        authorized_networks {
            value = "34.67.6.157"
        }

        authorized_networks {
            value = "34.67.234.134"
        }

        authorized_networks {
            value = "34.72.239.218"
        }
      }
    }

    deletion_protection  = "false"
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-my-profile%{random_suffix}"

	postgresql_profile {
		hostname = google_sql_database_instance.instance.public_ip_address
		username = google_sql_user.user.name
		password = google_sql_user.user.password
		database = google_sql_database.db.name
	}
	%{lifecycle_block}
}
`, context)
}

func testAccDatastreamConnectionProfile_mySQLUpdate(context map[string]interface{}, preventDestroy bool, password string) string {
	context["lifecycle_block"] = ""
	if preventDestroy {
		context["lifecycle_block"] = `
		lifecycle {
			prevent_destroy = true
		}`
	}

	context["password"] = password

	return acctest.Nprintf(`
resource "google_sql_database_instance" "mysql_instance" {
    name             = "tf-test-mysql-database-instance%{random_suffix}"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
      tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

      ip_configuration {

        // Datastream IPs will vary by region.
        authorized_networks {
            value = "34.71.242.81"
        }

        authorized_networks {
            value = "34.72.28.29"
        }

        authorized_networks {
            value = "34.67.6.157"
        }

        authorized_networks {
            value = "34.67.234.134"
        }

        authorized_networks {
            value = "34.72.239.218"
        }
      }
    }

    deletion_protection  = "false"
}

resource "google_sql_database" "mysql_db" {
    instance = google_sql_database_instance.mysql_instance.name
    name     = "db"
}

resource "google_sql_user" "mysql_user" {
    name = "user"
    instance = google_sql_database_instance.mysql_instance.name
    host     = "%"
    password = "%{password}"
}

resource "google_datastream_connection_profile" "mysql_con_profile" {
    display_name          = "Source connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-mysql-profile%{random_suffix}"

    mysql_profile {
		hostname = google_sql_database_instance.mysql_instance.public_ip_address
		username = google_sql_user.mysql_user.name
		password = google_sql_user.mysql_user.password
	}
	%{lifecycle_block}
}
`, context)
}

func testAccDatastreamConnectionProfile_sshKey_update(context map[string]interface{}, preventDestroy bool, private_key string, public_key string) string {
	context["lifecycle_block"] = ""
	if preventDestroy {
		context["lifecycle_block"] = `
        lifecycle {
            prevent_destroy = true
        }`
	}
	context["private_key"] = private_key
	context["public_key"] = public_key

	return acctest.Nprintf(`
resource "google_compute_network" "default" {
		name = "tf-test-datastream-ssh%{random_suffix}"
}

resource "google_sql_database_instance" "instance" {
    depends_on         = [google_compute_instance.default]
    name            	= "tf-test-my-database-instance%{random_suffix}"
    database_version	= "POSTGRES_14"
    region           	= "us-central1"
    settings {
        tier = "db-f1-micro"
        ip_configuration {
			ipv4_enabled = true

			authorized_networks {
				value = google_compute_instance.default.network_interface.0.access_config.0.nat_ip
			}
        }
    }
    
    deletion_protection  = "false"
}

resource "google_sql_database" "db" {
	depends_on = [google_sql_database_instance.instance]
	instance = google_sql_database_instance.instance.name
	name     = "db"
}

resource "google_sql_user" "user" {
	depends_on	= [google_sql_database_instance.instance]
	name		= "user"
	instance	= google_sql_database_instance.instance.name
	password	= "password%{random_suffix}"
}

resource "google_compute_instance" "default" {
	name         = "tf-test-instance-%{random_suffix}"
	machine_type = "e2-small"
	zone         = "us-central1-a"
	boot_disk {
		initialize_params {
			image = "debian-11-bullseye-v20241009"
		}
	}

	network_interface {
		network    = google_compute_network.default.name
		access_config {}
		}

	metadata = {
		"ssh-keys" = "user:%{public_key}"
	}

	metadata_startup_script = <<-EOT
	#!/bin/bash
	echo "Updating SSHD config for SSH forwarding..."

	# Backup sshd_config
	echo "AllowTcpForwarding yes" >> /etc/ssh/sshd_config
	echo "PasswordAuthentication no" >> /etc/ssh/sshd_config
	echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config
	echo "AuthorizedKeysFile .ssh/authorized_keys" >> /etc/ssh/sshd_config
	
	# Restart SSH service
	systemctl restart sshd
	EOT

	tags = ["ssh-host"]

	depends_on = [google_compute_firewall.ssh, google_compute_firewall.datastream_sql_access]

}

resource "time_sleep" "ssh_host_wait" {
	depends_on = [google_compute_instance.default]
	create_duration = "12m"
}

resource "google_compute_firewall" "ssh" {
	name 	= "tf-test-%{random_suffix}"
	network =  google_compute_network.default.name

	allow {
		protocol = "tcp"
		ports    = ["22"]
	}

	direction     = "INGRESS"
	priority      = 1000
	source_ranges = ["34.71.242.81", "34.72.28.29", "34.67.6.157", "34.67.234.134", "34.72.239.218"]

	target_tags = ["ssh-host"]
}

resource "google_compute_firewall" "datastream_sql_access" {
    name    	= "datastream-to-cloudsql-%{random_suffix}"
    network 	=  google_compute_network.default.name

    allow {
        protocol = "tcp"
        ports    = ["5432"]
    }

    direction     = "INGRESS"
    priority      = 1000
    source_ranges = ["34.71.242.81", "34.72.28.29", "34.67.6.157", "34.67.234.134", "34.72.239.218"]  #Datastream IPs

}

resource "google_datastream_connection_profile" "ssh_connectivity_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-pg-profile%{random_suffix}"

    postgresql_profile {
        hostname 			= google_sql_database_instance.instance.public_ip_address
        username 			= google_sql_user.user.name
        password 			= google_sql_user.user.password
        database 			= google_sql_database.db.name
        port 				= 5432
    }

    forward_ssh_connectivity {
        hostname 	= google_compute_instance.default.network_interface.0.access_config.0.nat_ip
        username 	= google_sql_user.user.name
        port    	= 22
        private_key 	= <<EOT
%{private_key}
EOT
	}

	depends_on = [time_sleep.ssh_host_wait]
	timeouts {
         create = "10m"
	}
    %{lifecycle_block}
}
`, context)
}
