package main

import (
  "strings"
	"regexp"
)

func validate(item jedi) bool {
  clearErrors()

  return validateName(item.Name) && validateEmail(item.Email)
}

func clearErrors()  {
  errors["name"] = ""
  errors["email"] = ""
}

func validateName(name string) bool {

  trimmedName := strings.Trim(name, " .")

  if len(trimmedName) == 0 {
    errors["name"] = "Name field is required"

    return false
  }

  reg := regexp.MustCompile("^[a-zA-Z. ]*$")
  res := reg.MatchString(name)

  if (res == false) {
    errors["name"] = "Only english letters, dots, spaces are allowed."

    return false
  }

  return true
}

func validateEmail(email string) bool {

  if len(email) == 0 {
    errors["email"] = "Email field is required"

    return false
  }

  reg := regexp.MustCompile("^[a-zA-Z0-9.-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9]{2,6}$")
  res := reg.MatchString(email)

  if (res == false) {
    errors["email"] = "Email format is incorrect."

    return false
  }

  for i := range jedis {
    if jedis[i].Email == email {
        errors["email"] = "Jedi with the same email is already exist!"

        return false
    }
  }

  return true
}
