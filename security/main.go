package security

import "github.com/keybase/go-keychain"

type Creds struct {
	Service string
	Account string
	Label   string
}

func (c *Creds) Set(username, password string) {
	item := keychain.NewGenericPassword(c.Service, c.Account, c.Label, []byte(username+"\n"+password), "")
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	err := keychain.AddItem(item)
	if err != nil {
		panic(err)
	}
}

func (c *Creds) Get() string {
	password, err := keychain.GetGenericPassword(c.Service, c.Account, c.Label, "")
	if err != nil {
		panic(err)
	}

	return string(password)
}
