package hello

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/mail"
    "strings"
)

func init() {
    http.HandleFunc("/", handler)
}

type EmailUser struct {
    Username    string
}


//pw: CoinsSose
var emailUsers = []string{"coins.sose15.stein@gmail.com", "coins.sose15.post@gmail.com", "coins.sose15.klebula@gmail.com", "coins.sose15.rossmehl@gmail.com","coins.sose15.klinger@gmail.com"}


func handler(w http.ResponseWriter, r *http.Request) {



    if err := r.ParseForm(); err != nil {
        return
    }

    var user = r.PostFormValue("user_id")

    var senderEmail = "coins.sose15.stein@gmail.com"

    if (user == "U04EE80D4") {
        senderEmail = "coins.sose15.klebula@gmail.com"
    }

    if (user == "U04ELLD9S") {
        senderEmail = "coins.sose15.rossmehl@gmail.com"
    }

    if (user == "U04EFCXQJ") {
        senderEmail = "coins.sose15.klinger@gmail.com"
    }

    if (user == "U04F7STS3") {
        senderEmail = "coins.sose15.post@gmail.com"
    }


    textMessage := r.PostFormValue("text")
    var mentions, notMentioned, newMailMessage = getMentionedAndNotMentionedAndText(textMessage)

    shortText := ""
    if (len(newMailMessage) > 50) {
        shortText = newMailMessage[0:50]
        shortText += "..."
    } else {
        shortText = newMailMessage
    }

    subject := "[SLACK]: " + shortText

    c := appengine.NewContext(r)



    if (len(mentions) == 0) {
        msg := &mail.Message{
            Sender:  senderEmail,
            To:      emailUsers,
            Cc: []string{"coinproject9@gmail.com"},
            Subject: subject,
            //Body:    textMessage,
            Body: newMailMessage,
        }
        if err := mail.Send(c, msg); err != nil {
            fmt.Fprint(w,err)
        }
    } else {
        notMentioned = append(notMentioned, "coinproject9@gmail.com")
        msg := &mail.Message{
            Sender:  senderEmail,
            To:      mentions,
            Cc: notMentioned,
            Subject: subject,
            //Body:    textMessage,
            Body: newMailMessage,
        }
        if err := mail.Send(c, msg); err != nil {
            fmt.Fprint(w,err)
        }
    }


}


func getMentionedAndNotMentionedAndText(text string) ([]string,[]string, string){
    var mentioned = []string{}
    var notMentioned = []string{}
    var newText = text

    if (strings.Contains(text, "<@U04EFCXQJ>")) {
        mentioned = append(mentioned, "coins.sose15.klinger@gmail.com")
        newText = strings.Replace(text, "<@U04EFCXQJ>", "@Daniel",-1)
    } else {
        notMentioned = append(notMentioned, "coins.sose15.klinger@gmail.com")
    }

    if (strings.Contains(text, "<@U04EE80D4>")) {
        mentioned = append(mentioned, "coins.sose15.klebula@gmail.com")
        newText = strings.Replace(text, "<@U04EE80D4>", "@Eike",-1)
    } else {
        notMentioned = append(notMentioned, "coins.sose15.klebula@gmail.com")
    }

    if (strings.Contains(text, "<@U04EE78B>")) {
        mentioned = append(mentioned, "coins.sose15.stein@gmail.com")
        newText = strings.Replace(text, "<@U04EE78B>", "@Marius",-1)
    } else {
        notMentioned = append(notMentioned, "coins.sose15.stein@gmail.com")
    }

    if (strings.Contains(text, "<@U04ELLD9S>")) {
        mentioned = append(mentioned, "coins.sose15.rossmehl@gmail.com")
        newText = strings.Replace(text, "<@U04ELLD9S>", "@Max",-1)
    } else {
        notMentioned = append(notMentioned, "coins.sose15.rossmehl@gmail.com")
    }

    if (strings.Contains(text, "<@U04F7STS3>")) {
        mentioned = append(mentioned, "coins.sose15.post@gmail.com")
        newText = strings.Replace(text, "<@U04F7STS3>", "@Alexander",-1)
    } else {
        notMentioned = append(notMentioned, "coins.sose15.post@gmail.com")
    }

    return mentioned, notMentioned, newText;
}
