Terraform Controller
===
Generate and manage AWS resources from Kubernetes CRD's (Using terraform)

Context
===
Traditionally managing terraform repos is abstracted away from your development teams (in the case of a mono repo)
and requires all teams working with other infrastructure resources to have development knowledge of terraform.

Problem Statement
===
The problem is that its just hard for a software developer to manage all infrastructure dependencies to their applications in a cloud environment (AWS/GCP.. etc). Terraform was created to solve this problem, but now we need to learn to use terraform, in a consistent manner and try not to manually screw it up by misusing names and managing shared/global state.

Solution
===
Abstract all that away inside a neat little operator with a bunch of CRD's that describe what your application needs, and can be deployed alongside the application in a nice simple way (using helm or similar).

Usage
===

TODO
