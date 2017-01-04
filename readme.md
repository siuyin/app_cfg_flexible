# Flexible Manager/Supervisor program configuration
1. Configuration file (yaml)
  1. By Board (NATS topic) type
    1. By Message detail

Note "gopkg.in/yaml.v2" idiosyncrasies. Structs in Go have to be
exported (Uppercase) but yaml keys are lowercase.
Use Go struct tags to make consistent. 
