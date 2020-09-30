package main

import (
	"fmt"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	//adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	//usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	msg := "Welcome! I'm Porus the bot. \n\n My mentor is still teahcing me things.\n For now, here is your allowed commands \n\n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
	
}

func (a *application) p4d9Handler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ps4d9msg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ps4d9i := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ps4d9p := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	ps4d9o := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ps4d9h := fmt.Sprintf("\n { ---Anonymous Phone Packages--- } \n%s %s %s", ps4d9i, ps4d9p, ps4d9o)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ps4d9h)
	a.client.SendMessage(m.Chat.ID, ps4d9msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) p3z2Handler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ps3z2msg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ps3z2i := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ps3z2p := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	ps3z2o := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ps3z2h := fmt.Sprintf("\n { ---Data Removal Services--- } \n%s %s %s", ps3z2i, ps3z2p, ps3z2o)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ps3z2h)
	a.client.SendMessage(m.Chat.ID, ps3z2msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) p7e1Handler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ps7e1msg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ps7e1i := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ps7e1p := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	ps7e1o := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ps7e1h := fmt.Sprintf("\n { ---Consultations--- } \n%s %s %s", ps7e1i, ps7e1p, ps7e1o)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ps7e1h)
	a.client.SendMessage(m.Chat.ID, ps7e1msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) p8j0Handler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ps8j0msg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ps8j0i := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ps8j0p := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	ps8j0o := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ps8j0h := fmt.Sprintf("\n { ---Private Investigation--- } \n%s %s %s", ps8j0i, ps8j0p, ps8j0o)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ps8j0h)
	a.client.SendMessage(m.Chat.ID, ps8j0msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}






func (a *application) porusHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	msg := "[InternalMem@Porus]:\n I was named Porus for my soon to come resourcefulness or expediency. \n I will: \n> Securely handle orders available for automation from the payment to delivery \n>Generate business addresses & shipping labels for you \n>Provide Anonymous Tracking Check!\n> Provide automated recon on targets(data; public/private records, reverse-search(license plates, emails, usernames, etc), much more!) "
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency	
}
func (a *application) psHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	Cashapph := fmt.Sprintf("\n { ---Propducts & Services--- } \n [+] Commands: \n Use /ps1y4 For Cashapp Account Information& Ordering \n\n Use /ps4d9 For Anonymous Phone Packages Infromation & Ordering\n\n Use /ps3z2 For Data Removal Servies Info & Ordering \n\n Use /ps7e1 For Consultation Information & Scheduling")
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, Cashapph)
	// a.client.SendMessage(m.Chat.ID, Cashappmsg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) pauyHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ps1y4msg := fmt.Sprintf("\n Order: *CASHAPP ACCOUNT*\n Quantity: *1or2or5*\n Payment Method: *BTCorCASHAPP\n Address: *For CshAPP Card*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ps1y4i := fmt.Sprintf(">Safest & Freshest Accounts On The Market! \n>We NEVER sell you stolen or hot accounts.\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ps1y4p := fmt.Sprintf("[+] Current Price Is: \n> Single Purchase: $500 \n> 2 Cashapp accounts: $900\n> 5 Cashapp Accounts: $1500  \n(IF MORE THAN THIS CONTINUE)\n")
	ps1y4o := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	pauyh := fmt.Sprintf("\n { ---Anonymous Cashapp Accounts--- } \n\n %s %s %s", ps1y4i, ps1y4p, ps1y4o)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, pauyh)
	a.client.SendMessage(m.Chat.ID, ps1y4msg)
	// a.client.SendMessage(m.Chat.ID, ps1y4msg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}




// Handle the /play command here
func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an option:", tbot.OptInlineKeyboardMarkup(buttons))
}

// Handle the /score command here
func (a *application) scoreHandler(m *tbot.Message) {
	msg := fmt.Sprintf("Scores:\nWins: %v\nDraws: %v\nLosses: %v", a.wins, a.draws, a.losses)
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /reset command here
func (a *application) resetHandler(m *tbot.Message) {
	a.wins, a.draws, a.losses = 0, 0, 0
	msg := "Scores have been reset to 0."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle buttton presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := a.draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
