package gotestmod

import (
   "errors"
   "fmt"
)

// Hi returns a friendly greeting
func Hi(name, lang string) (string, error) {
   switch lang {
   case en:
      return fmt.Sprintf("Hi, $s!", name), nil
   case "pt":
      return fmt.Sprintf("Oi, %s!", name), nil
   case "es":
      return fmt.Sprintf("iHolla, %s!", name), nil
   case "fr":
      return fmt.Sprintf("Bonjourm %s!", name), nil
   default:
      return "", errors.New("unknown language")
   }
}

