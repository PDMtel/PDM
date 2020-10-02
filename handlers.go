package main

import (
	"fmt"
	"strings"
	"net"
	"sort"
	
	"github.com/yanzay/tbot/v2"
)

func worker(hosty, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", hosty, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	//
	text := strings.TrimPrefix(m.Text, "/start ")
	adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", m.CallbackQuery.Data)
	msg := "Welcome! I'm Porus the bot. \n\n My mentor is still teahcing me things.\n For now, here is your allowed commands \n\n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /late To report a late/missing order\n4. Use /supp For Support An Your Old/Current Order\n\nUse /methods For avalible sacue for sell\n\nUse /legalflips For legal investments\nUse /freetip for a free random tip "
	msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, text) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
	
}
				       
				
func (a *application) pscanHandler(m *tbot.Message) {
	text := strings.TrimPrefix(m.Text, "/pscan ")
	hosty := text
		//
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	
	for i := 0;i < cap(ports); i++ {
		go worker(hosty, ports, results)
	}
	
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		xx := fmt.Sprintf("%d open\n", port)
		a.client.SendMessage(m.Chat.ID, xx)
	}
	
	
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	msg := "[InternalMem@Porus]:\n I was named Porus for my soon to come resourcefulness or expediency. \n I will: \n> Securely handle orders available for automation from the payment to delivery \n>Generate business addresses & shipping labels for you \n>Provide Anonymous Tracking Check!\n> Provide automated recon on targets(data; public/private records, reverse-search(license plates, emails, usernames, etc), much more!) "
	msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, text) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	a.client.SendMessage(adminID, msgadmin)
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

func (a *application) lateHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	msg := "[InternalMem@Porus]:\n If your reporting a late order: \n\n I apologize, I will give you a gift for your patience and understanding. I admit my communication, organization, delivery times are still improving. That is exactly why I devleop things like this to save me time. I am not a 'snap trapper', but an enthusiast of technnology, security & privacy. I seen how vulnerable the game was and needed me. I joined mainly to help you & I enjoy it.\n\n\n To report a late order: \nCopy The Next Message As a Template\n"
	msgh := "Order: LATE !!\nOrginal Payment Method: *YourFormOfPay*\nPaymentMethod ID: *$username*\nOrder PlacedDate: *DoNotExaggerate\nYour Expecting: *UrOrderHere\nAdditional Info:*anything*\nProof:\n\n"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	a.client.SendMessage(m.Chat.ID, msgh)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency	
}
func (a *application) suppHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	msg := "[InternalMem@Porus]:\n If You Need Order Support: \n\n I will get back to you soon as I am avalible. All inquires go in order. \n\n Copy And {Aste the next message as a template. "
	msgh := "Order: SUPPORT !!\nOrginal Payment Method: *YourFormOfPay*\nPaymentMethod ID: *$username*\nNeed Help With: *UrOrderHere\nAdditional Info:*anything*\nProof:\n\n"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, msg)
	a.client.SendMessage(m.Chat.ID, msgh)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency	
}

func (a *application) psHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	Cashapph := fmt.Sprintf("\n { ---Products & Services--- } \n [+] Commands: \n Use /pauy For Cashapp Account Information& Ordering \n\n Use /ppp For Anonymous Phone Packages Infromation & Ordering\n\n Use /pdrs For Data Removal Servies Info & Ordering \n\n Use /pc For Consultation Information & Scheduling\nUse /ppi For Private Investigation Services")
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, Cashapph)
	// a.client.SendMessage(m.Chat.ID, Cashappmsg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}

func (a *application) pppHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	pppmsg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	pppi := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	pppp := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each) \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	pppo := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ppph := fmt.Sprintf("\n { ---Anonymous Phone Packages--- } \n%s %s %s", pppi, pppp, pppo)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ppph)
	a.client.SendMessage(m.Chat.ID, pppmsg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) pdrsHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	pdrsmsg := fmt.Sprintf("\n *")
	
	pdrsi := fmt.Sprintf(">REMOVE YOUR FOOTPRINT OFF THE WEB.\n>Today's age, we don't truly have privacy until We change\n>We make it easy for you.\n>We give you an asessment & full report on data removal done on you\n\n")
	pdrsp := fmt.Sprintf("[+] The Current Packages Are: \n {+()}\n> \n {+:}\n> \n")
	pdrso := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n ")
	pdrsh := fmt.Sprintf("\n { ---Data Removal Services--- } \n%s %s %s", pdrsi, pdrsp, pdrso)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, pdrsh)
	a.client.SendMessage(m.Chat.ID, pdrsmsg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) pcHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	pcmsg := fmt.Sprintf("\n Order: PHONE PACKAGE \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	pci := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	pcp := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	pco := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	pch := fmt.Sprintf("\n { ---Consultations--- } \n%s %s %s", pci, pcp, pco)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, pch)
	a.client.SendMessage(m.Chat.ID, pcmsg)
	// a.client.SendMessage(adminID, msgadmin) //notify me of commadns sent
	// resourcefulness or expediency
}
func (a *application) ppiHandler(m *tbot.Message) {
	// adminID := fmt.Sprintf("1331473188") //Notify me of commands sent
	// usercomm := fmt.Sprintf("%s", tbot.CallbackQuery.Data)
	ppimsg := fmt.Sprintf("\n Order: PII \n Which Package?: *Use The package title*\nQuantity: *1to100*\n Payment Method: *BTCorCASHAPP*\n Address: *Any*\n Custom Name: *AnyNameYouWant*\n Referral: *BLANKorFILL*\n Addtional Info: *Anything I should Know*")
	
	ppii := fmt.Sprintf(">ALL PHONES CAN COME WITH CASHAPP & YEAR OF SERVICE! \n> All Sim Cards In Our Phones Can Be Used In Any Device Too!\nWe NEVER sell you stolen or hot products,\n>ALL cashapp accounts are level2 verified via vulnerbility(non-detection) & ready to recieve 25k a week & cash out.\n>WITH A 30-DAY GUARANTEE WARRANTY\n\n")
	ppip := fmt.Sprintf("[+] The Current Prices Are: \n {+Preloaded Sec/Privacy Phones(with year of service) No Cashapp}\n> MotorolaE Single Purchase: $250(Each) \n> 2-5 Phones: $225 Each \n> 5-20 Phones: $200(Each)  \n {+Basic Packages:MotorolaE,Cshapp,CshCard,PreLoaded-Sec/Privacy}\n> Single Purchase: $500 \n> 2: $900\n> 5: $1800 \n(IF MORE THAN THIS CONTINUE)\n")
	ppio := fmt.Sprintf("I am currently not ready to automate this order yet! \n\n [+]--TO PLACE A ORDER:\n {FOLLOW EVERY STEP}\n, \n Message @XskotosX On Telegram This: \n YES IT HAS TO LOOK LIKE THIS.")
	
	ppih := fmt.Sprintf("\n { ---Private Investigation--- } \n%s %s %s", ppii, ppip, ppio)
	
	
	//msg := "Welcome! I'm Porus the bot. My mentor is still teahcing me things. For now, here is your allowed commands \n[+] Commands:\n1. Use /porus to know more about me(porus) and why I was developed \n2. Use /ps to know avalaible products & services with order instructions \n3. Use /LATE for LATE orders and support instructions"
	//msgadmin := fmt.Sprintf("/start command initiated by Chat ID %s:%d:%s", m.Chat.ID, m.MessageID, usercomm) //Notify me of commands sent
	a.client.SendMessage(m.Chat.ID, ppih)
	a.client.SendMessage(m.Chat.ID, ppimsg)
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
