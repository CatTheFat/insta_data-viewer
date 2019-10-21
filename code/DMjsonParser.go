package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

//ErrHandle : prints an error if there is one
func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)

		_, _, line, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("Error check called from #%d\n", line)
		}
	}
}

//Message : data included in every message
type Message struct {
	Sender            string `json:"sender"`
	Timestamp         string `json:"created_at"`
	Text              string `json:"text"`
	MediaShareOwner   string `json:"media_owner"`
	MediaShareCaption string `json:"media_share_caption"`
	MediaShareURL     string `json:"media_share_url"`
	MediaSent         string `json:"media"`
	MediaSent1        string `json:"media_url"`
	Heart             string `json:"heart"`
	Action            string `json:"action"`
	VideoCallAction   string `json:"video_call_action"`
	StoryShare        string `json:"story_share"`
	Audio             string `json:"live_video_invite"`
}

//MsgBlock : each conversation is housed in it
type MsgBlock struct {
	Participants []string
	Conversation []Message
}

//--------------------------------------------//

var jsonData []MsgBlock
var master string

func openf() {
	var fileLoc string

	fmt.Println(`Enter the "messages.json" file location`)
	fmt.Scanln(&fileLoc)

	fmt.Println(`Enter the owner of the file (@tag)`)
	fmt.Scanln(&master)

	jsonRawBytes, fopenErr := ioutil.ReadFile("C:\\MaPath\\files\\messages.json")

	//---ReadFile error handling---//
	ErrHandle(fopenErr)
	//-----------------------------//

	jparseErr := json.Unmarshal(jsonRawBytes, &jsonData)

	//---Unmarshal error handling---//
	ErrHandle(jparseErr)
	//------------------------------//

}

func list() {
	var name string
	isGroup := false

	fmt.Printf("\nShowing %v conversation(s):\n", len(jsonData))

	//jsonDataList := *jsonDatap

	for i, block := range jsonData {

		for _, participant := range block.Participants {
			if participant == master {
				continue
			}
			if len(block.Participants) >= 3 {
				isGroup = true
			}
			name = participant
		}
		if !isGroup {
			fmt.Printf("[%v] - %v\n", i, name)
		} else {
			fmt.Printf("[%v] - group: participants %+v\n", i, block.Participants)
			isGroup = false
		}
	}

	fmt.Println("")
}

func createDir(standalone bool) (mkDir string) {
	var dirPath string
	var mkDirSTRp *string

	fmt.Println("Enter the folder location:")
	fmt.Scanln(&dirPath)

	if standalone {

		mkDirSTR := dirPath + `\` + master
		mkDirSTRp = &mkDirSTR

	} else {

		mkDirSTR := dirPath + `\` + master + `\messages`
		mkDirSTRp = &mkDirSTR

	}

	mkdirErr := os.MkdirAll(*mkDirSTRp, 0777)

	ErrHandle(mkdirErr)

	return *mkDirSTRp
}

func createFile(mainDir string, name string) (fileP *os.File) {
	fileDir := mainDir + `\` + name + ".html"

	//Creates a file with the name of the main participant
	expFile, crFileErr := os.OpenFile(fileDir, os.O_CREATE|os.O_WRONLY, 0777)

	//Handles any file creation errors
	ErrHandle(crFileErr)

	return expFile
}

func getMSGtext(msgOBJ Message) string {
	switch {
	case msgOBJ.Text != "":
		return msgOBJ.Text
	case msgOBJ.Heart != "":
		return msgOBJ.Heart
	case msgOBJ.Action != "":
		return msgOBJ.Action
	case msgOBJ.VideoCallAction != "":
		return msgOBJ.VideoCallAction
	}

	return ""
}

func exportConv(id int, standalone bool, mainPath string) {
	var name string
	var numName string
	var standMainPATH string
	var usedPath *string

	//Creates the needed directories if not called by the export loop
	if standalone {
		standMainPATH = createDir(true)
	}

	//If the conversation is not a group
	if len(jsonData[id].Participants) < 3 {

		//Gets the name of the slave
		for _, participant := range jsonData[id].Participants {
			if participant == master {
				continue
			}
			name = participant
			numName = fmt.Sprintf("%v%v", id, participant)
		}

		//Sets the path to use for file creation
		if standalone {

			usedPath = &standMainPATH

		} else {

			usedPath = &mainPath

		}

		expFile := createFile(*usedPath, numName)

		//Writes the prefix to the file
		_, wrF0err := expFile.Write([]byte(`<!DOCTYPE html>
		<html>
		<head>
		<meta http-equiv=Content-Type content="text/html; charset=UTF-8, user-scalable=yes" />
		<style type=text/css>body{background:#fff;color:#1c1e21;direction:ltr;line-height:1.34;margin:0;padding:0;unicode-bidi:embed}body,button,input,label,select,td,textarea{font-family:Helvetica,Arial,sans-serif;font-size:12px}h1,h2,h3,h4,h5,h6{color:#1c1e21;font-size:13px;font-weight:600;margin:0;padding:0}h1{font-size:14px}h4,h5,h6{font-size:12px}b,strong{font-weight:600}a{color:#385898;cursor:pointer;text-decoration:none}button{margin:0}a:hover{text-decoration:underline}img{border:0}td,td.label{text-align:left}dd{color:#000}dt{color:#606770}ul{list-style-type:none;margin:0;padding:0}abbr{border-bottom:0;text-decoration:none}hr{background:#dadde1;border-width:0;color:#dadde1;height:1px}form{margin:0;padding:0}label{color:#606770;cursor:default;font-weight:600;vertical-align:middle}label input{font-weight:normal}textarea,.inputtext,.inputpassword{border:1px solid #ccd0d5;border-radius:0;margin:0;padding:3px}textarea{max-width:100%}select{border:1px solid #ccd0d5;padding:2px}input,select,textarea{background-color:#fff;color:#1c1e21}.inputtext,.inputpassword{padding-bottom:4px}.inputtext:invalid,.inputpassword:invalid{box-shadow:none}.inputradio{margin:0 5px 0 0;padding:0;vertical-align:middle}.inputcheckbox{border:0;vertical-align:middle}.inputbutton,.inputsubmit{background-color:#4267b2;border-color:#dadde1 #0e1f5b #0e1f5b #d9dfea;border-style:solid;border-width:1px;color:#fff;padding:2px 15px 3px 15px;text-align:center}.inputsubmit_disabled{background-color:#999;border-bottom:1px solid #000;border-right:1px solid #666;color:#fff}.inputaux{background:#ebedf0;border-color:#ebedf0 #666 #666 #e7e7e7;color:#000}.inputaux_disabled{color:#999}.inputsearch{background:#fff url(https://static.xx.fbcdn.net/rsrc.php/v3/yV/r/IJYgcESal33.png) no-repeat left 4px;padding-left:17px}.clearfix:after{clear:both;content:'.';display:block;font-size:0;height:0;line-height:0;visibility:hidden}.clearfix{zoom:1}.datawrap{word-wrap:break-word}.word_break{display:inline-block}.flexchildwrap{min-width:0;word-wrap:break-word}.ellipsis{overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.aero{opacity:.5}.column{float:left}.center{margin-left:auto;margin-right:auto}#facebook .hidden_elem{display:none!important}#facebook .invisible_elem{visibility:hidden}#facebook .accessible_elem{clip:rect(1px,1px,1px,1px);height:1px;overflow:hidden;position:absolute;white-space:nowrap;width:1px}.direction_ltr{direction:ltr}.direction_rtl{direction:rtl}.text_align_ltr{text-align:left}.text_align_rtl{text-align:right}body{overflow-y:scroll}.mini_iframe,.serverfbml_iframe{overflow-y:visible}.auto_resize_iframe{height:auto;overflow:hidden}.pipe{color:gray;padding:0 3px}#content{margin:0;outline:0;padding:0;width:auto}.profile #content,.home #content,.search #content{min-height:600px}.UIStandardFrame_Container{margin:0 auto;padding-top:20px;width:960px}.UIStandardFrame_Content{float:left;margin:0;padding:0;width:760px}.UIStandardFrame_SidebarAds{float:right;margin:0;padding:0;width:200px;word-wrap:break-word}.UIFullPage_Container{margin:0 auto;padding:20px 12px 0;width:940px}.empty_message{background:#f5f6f7;font-size:13px;line-height:17px;padding:20px 20px 50px;text-align:center}.see_all{text-align:right}.standard_status_element{visibility:hidden}.standard_status_element.async_saving{visibility:visible}img.tracking_pixel{height:1px;position:absolute;visibility:hidden;width:1px}#globalContainer{margin:0 auto;position:relative;zoom:1}.fbx #globalContainer{width:981px}.sidebarMode #globalContainer{padding-right:205px}.fbx #tab_canvas>div{padding-top:0}.fb_content{min-height:612px;padding-bottom:20px}.fbx .fb_content{padding-bottom:0}.skipto{display:none}.home .skipto{display:block}._li._li._li{overflow:initial}._72b0{position:relative;z-index:0}._5vb_ #pageFooter{display:none}html body._5vb_ #globalContainer{width:976px}._5vb_.hasLeftCol #headerArea{margin:0;padding-top:0;width:786px}._5vb_,._5vb_ #contentCol{background-color:#e9ebee;color:#1d2129}html ._5vb_.hasLeftCol #contentCol{border-left:0;margin-left:172px;padding-left:11px;padding-top:11px}._5vb_.hasLeftCol #topNav{border-left:0;margin-left:172px;padding:11px 7px 0 11px}._5vb_.hasLeftCol #topNav~#contentCol{padding-top:0}._5vb_.hasLeftCol #leftCol{padding-left:8px;padding-top:12px;width:164px}._5vb_.hasLeftCol #mainContainer{border-right:0;margin-left:0}._5vb_.hasLeftCol #pageFooter{background:0}html ._5vb_._5vb_.hasLeftCol div#contentArea{padding-left:0;padding-right:10px;width:786px}html ._5vb_._5vb_.hasLeftCol .hasRightCol div#contentArea{width:496px}._5vb_.hasLeftCol ._5r-_ div#rightCol{padding:0 7px 0 0;width:280px}._2yq #globalContainer{width:1012px!important}._2yq #headerArea{float:none!important;padding:0 0 12px 0!important;width:auto!important}._2yq #contentArea{margin-right:0;padding:0!important}._2yq #leftCol,._2yq #contentCol{padding:0!important}._2yq #rightCol{float:left;margin-top:0;padding:0!important}._2yq .groupJumpLayout{margin-top:-12px}._2yq .loggedout_menubar_container{min-width:1014px}._4yic{margin:auto;font-size:13px}._3a_u{margin:0 auto;width:50%}._4t5n{float:none;margin-bottom:15px;position:relative;word-break:break-word;z-index:0}._4t5o{clear:both;color:#7f7f7f;font-size:14px;margin-bottom:20px;margin-top:10px;text-align:center}._3b0a{position:relative;z-index:100}._3b0b{background:#fff;border-radius:3px;display:flex;flex-direction:row;padding:15px}._3z-t{border-radius:50%;height:16px;padding:4px;width:16px}._3b0c{display:flex;flex-direction:column;justify-content:center;margin-left:8px}._3b0d{color:#1d2129;font-size:14px;font-weight:bold;line-height:18px;margin-bottom:3px}._3b0e{color:#90949c;font-size:14px;line-height:16px}._218o{display:flex;justify-content:space-between;margin:0 auto;margin-left:60%;width:auto}._72m4{font-size:12px;margin-top:4px}._5aj7{display:flex}._5aj7 ._4bl7{float:none}._5aj7 ._4bl9{flex:1 0 0}._ikh ._4bl7{float:left;min-height:1px}._4bl7,._4bl9{word-wrap:break-word}._4bl9{overflow:hidden}._21dp{position:relative;z-index:301}._2t-8._2t-8{font-family:Helvetica,Arial,sans-serif}._2t-8{height:43px;min-width:100%}._2t-a{height:42px;position:relative;width:100%}._2s1y{box-sizing:border-box;height:43px}._50ti{position:fixed;top:0}.hasDemoBar ._50ti{top:60px}._33rf{min-width:981px}._2yq ._33rf,._2xk0 ._33rf{min-width:1014px}._50tj{box-sizing:border-box;padding-right:0}.sidebarMode ._50tj{padding-right:205px}._4pmj{box-sizing:border-box;padding:0 16px}._2t-d{margin:auto;padding:0 8px}._2t-a{display:flex;justify-content:space-between}._2t-e,._2t-f{display:flex}._2t-e{flex:1 1 auto;justify-content:flex-start}._2t-f{flex:0 0 auto;justify-content:flex-end;margin-left:8px}._4kny{float:left}._50tm{width:100%}._2s24{margin-left:1px;position:relative}._h2p ._2s24{margin-left:0}._cy6{display:inline-block;padding:0 9px;vertical-align:top}._h2p ._cy6{padding:0 5px 0 4px}._cy6:first-child{padding-left:0}._h2p ._cy6:first-child{padding:0}._cy6:last-child{padding-right:0}._cy7{margin:7px 0 8px 0}._h2p ._cy7{margin-right:1px}._2s25{background-color:transparent;color:inherit;display:inline-block;font-size:12px;font-weight:bold;height:27px;line-height:28px;padding:0 10px 1px;position:relative;text-decoration:none;vertical-align:top;white-space:nowrap}.segoe ._2s25{font-weight:600}._h2p ._2s25{padding:0 12px 1px}._h2p ._cy6 ._4kny:last-child ._2s25{padding-right:11px}._2s25:hover,._2s25:focus,._2s25:active{border-radius:2px;color:inherit;outline:0;text-decoration:none;z-index:1}.openToggler ._2s25:hover,.openToggler ._2s25:focus,.openToggler ._2s25:active{background:transparent}._4kny ._585-{margin-left:0;min-width:144px;width:100%}._4kny .__wu ._539-.roundedBox{margin-left:0}._4kny ._4962{float:none;margin:5px 0 6px 0;position:relative}._h2p ._4kny ._4962{margin:5px 0 6px}._3x1p{height:100%;overflow:hidden;position:absolute;width:100%}._63i8{display:block;left:0;opacity:0;position:absolute;top:0;transition:opacity 200ms;width:100vw;z-index:2}._4yim,._4yin{display:block;left:0;position:absolute;top:0;transform-origin:left}._4yin{transform-origin:right}._4yio,._4yip{display:inline-block;height:2px;left:0;min-width:12px;position:absolute;top:0;transform-origin:left}._4yip{transform-origin:right}._2t-8.indeterminateBarTransition.transitioning ._63i8{opacity:1}._2t-8.indeterminateBarTransition.transitioning ._4yim{animation:indeterminateBarTransitionTranslate-left 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yin{animation:indeterminateBarTransitionTranslate-right 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yio{animation:indeterminateBarTransitionWidth-left 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yip{animation:indeterminateBarTransitionWidth-right 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.finishing ._63i8{opacity:0}@keyframes indeterminateBarTransitionTranslate-left{0%{animation-timing-function:ease-in;transform:translateX(0)}25%{animation-timing-function:ease-out;transform:translateX(25vw)}50%{animation-timing-function:ease-in;transform:translateX(calc(100vw - 12px))}100%{transform:translateX(calc(100vw - 12px))}}@keyframes indeterminateBarTransitionTranslate-right{0%{transform:translateX(0)}51%{animation-timing-function:ease-in;transform:translateX(0)}75%{animation-timing-function:ease-out;transform:translateX(calc(-25vw))}100%{animation-timing-function:ease-in;transform:translateX(calc(-100vw+12px))}}@keyframes indeterminateBarTransitionWidth-left{0%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}25%{animation-timing-function:ease-out;transform:scaleX(50)}50%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}51%{opacity:0}100%{opacity:0}}@keyframes indeterminateBarTransitionWidth-right{0%{opacity:0}25%{opacity:0}50%
		{animation-timing-function:ease-in;opacity:0;transform:scaleX(1)}51%{opacity:1}75%{animation-timing-function:ease-out;opacity:1;transform:scaleX(50)}100%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}}._2s1x ._2s1y{background-color:#fff;color:#fff;box-shadow:0 4px 5px 0 rgba(179,179,179,1)}._2s1x ._2s24::before{background:rgba(0,0,0,.1)}._2s1x ._2s25:hover,._2s1x ._2s25:focus,._2s1x ._2s25:active{background:rgba(0,0,0,.1);color:inherit}._2s1x.transitioning ._3fju{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yv/r/z1PAxf53vph.gif);height:100%;width:100%}._2s1x.transitioning ._3fjv{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/y5/r/OzkCShPcfVN.gif);height:100%;width:100%}._2s1x.transitioning ._3fjx{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yH/r/YIs0iw6cHAa.gif);height:100%;width:100%}._2s1x ._3b33{background:#fff;height:100%;opacity:0;position:absolute;width:100%}._2s1x.pulseTransition.transitioning ._3b33{animation:pulse-loading 800ms cubic-bezier(.455,.03,.515,.955) infinite alternate;animation-timing-function:steps(8,end)}@keyframes pulse-loading{0%{opacity:0}100%{opacity:.15}}._2s1x ._3b34{height:100%;max-width:1016px;position:absolute;width:100%}._2s1x.shimmerTransition.transitioning ._3b34{animation:shimmer-loading 1600ms cubic-bezier(.455,.03,.515,.955) infinite;animation-timing-function:steps(16,end);background:linear-gradient(to right,rgba(66,103,178,0),#577fbc,rgba(66,103,178,0));background-size:1016px auto}@keyframes shimmer-loading{0%{transform:translateX(-1016px)}100%{transform:translateX(calc(100vw+1016px))}}._2s1x ._63tk{background-color:#fff}._19ea{margin:7px 0;margin-left:-2px;margin-right:5px}._19eb{display:inline-block;outline:0;padding:2px;position:relative}._7tp1{display:inline-block;height:20px;outline:0;padding:4px;position:relative}._19eb:hover,._19eb:focus,._19eb:active,._7tp1:hover,._7tp1:focus,._7tp1:active{background-color:#365899;background-color:rgba(0,0,0,.1);border-radius:3px}._2md{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yr/r/Xo5dkI2ODGj.png);background-repeat:no-repeat;background-size:auto;background-position:-25px -116px;display:block;height:24px;outline:0;overflow:hidden;text-indent:-999px;white-space:nowrap;width:24px}._7tp2{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yr/r/Xo5dkI2ODGj.png);background-repeat:no-repeat;background-size:auto;background-position:0 0;display:block}._7ql{border-radius:2px;display:inline;margin:2px 6px 2px -8px;vertical-align:inherit}._h2p ._7ql{margin-left:-10px}._1k67 ._2s25{position:relative}._1k67._d0b ._2s25{padding-right:0}._1k67 ._1vp5 .img{transform:translateY(2px)}._1k67._d0b._5-y2 ._2s25{padding-right:6px}._2qgu._2qgu{border-radius:50%;overflow:hidden}._2s25._2s25._606w._606w:after,._606w:after{border-radius:50%}._605a .fbxWelcomeBoxBlock:after{border-radius:50%}._1qv9{align-items:center;display:flex;flex-direction:row}._rv{height:100px;width:100px}._rw{height:50px;width:50px}._s0:only-child{display:block}._3tm9{height:14px;width:14px}._54rv{height:16px;width:16px}._3qxe{height:19px;width:19px}._1m6h{height:24px;width:24px}._3d80{height:28px;width:28px}._54ru{height:32px;width:32px}._tzw{height:40px;width:40px}._54rt{height:48px;width:48px}._54rs{height:56px;width:56px}._1m9m{height:64px;width:64px}._ry{height:24px;width:24px}._4jnw{margin:0}._3-8h{margin:4px}._3-8i{margin:8px}._3-8j{margin:12px}._3-8k{margin:16px}._3-8l{margin:20px}._2-5b{margin:24px}._1kbd{margin-bottom:0;margin-top:0}._3-8m{margin-bottom:4px;margin-top:4px}._3-8n{margin-bottom:8px;margin-top:8px}._3-8o{margin-bottom:12px;margin-top:12px}._3-8p{margin-bottom:16px;margin-top:16px}._3-8q{margin-bottom:20px;margin-top:20px}._2-ox{margin-bottom:24px;margin-top:24px}._1a4i{margin-left:0;margin-right:0}._3-8r{margin-left:4px;margin-right:4px}._3-8s{margin-left:8px;margin-right:8px}._3-8t{margin-left:12px;margin-right:12px}._3-8u{margin-left:16px;margin-right:16px}._3-8v{margin-left:20px;margin-right:20px}._6bu9{margin-left:24px;margin-right:24px}._5soe{margin-top:0}._3-8w{margin-top:4px}._3-8x{margin-top:8px}._3-8y{margin-top:12px}._3-8z{margin-top:16px}._3-8-{margin-top:20px}._4aws{margin-top:24px}._2-jz{margin-right:0}._3-8_{margin-right:4px}._3-90{margin-right:8px}._3-91{margin-right:12px}._3-92{margin-right:16px}._3-93{margin-right:20px}._y8t{margin-right:24px}._5emk{margin-bottom:0}._3-94{margin-bottom:4px}._3-95{margin-bottom:8px}._3-96{margin-bottom:12px}._3-97{margin-bottom:16px}._3-98{margin-bottom:20px}._20nr{margin-bottom:24px}._av_{margin-left:0}._3-99{margin-left:4px}._3-9a{margin-left:8px}._3-9b{margin-left:12px}._3-9c{margin-left:16px}._3-9d{margin-left:20px}._4m0t{margin-left:24px}._2lej{border-radius:20px}._2lek{color:#1d2129;font-size:14px;font-weight:bold;line-height:18px}._2lak{color:#fff;font-size:14px;font-weight:bold;line-height:18px}._2lel{border-bottom:1px solid #dadde1}._2lem,._2lem a{color:#8d949e;font-size:13px;line-height:16px}._2let{color:#1d2129;font-size:18px;line-height:22px}._2lat{color:#fff;font-size:18px;line-height:22px}._tqp{color:gray;font-size:13px}._4mp8{font-weight:bold}._4nkx,._3ttj{font-size:14px;line-height:2;text-align:left}._4nkx tbody tr th{padding:5px 5px;text-align:left;vertical-align:top;width:150px}._2yuc{max-width:100%}._3hls{font-size:14px;font-weight:bold}._2oao{color:#90949c;font-size:13px;font-weight:bold;line-height:20px;width:100px}._23bw{font-size:13px}._6udd{word-break:break-all}._8tm{padding:0}._2phz{padding:4px}._2ph-{padding:8px}._2ph_{padding:12px}._2pi0{padding:16px}._2pi1{padding:20px}._40c7{padding:24px}._2o1j{padding:36px}._6buq{padding-bottom:0;padding-top:0}._2pi2{padding-bottom:4px;padding-top:4px}._2pi3{padding-bottom:8px;padding-top:8px}._2pi4{padding-bottom:12px;padding-top:12px}._2pi5{padding-bottom:16px;padding-top:16px}._2pi6{padding-bottom:20px;padding-top:20px}._2o1k{padding-bottom:24px;padding-top:24px}._2o1l{padding-bottom:36px;padding-top:36px}._6bua{padding-left:0;padding-right:0}._2pi7{padding-left:4px;padding-right:4px}._2pi8{padding-left:8px;padding-right:8px}._2pi9{padding-left:12px;padding-right:12px}._2pia{padding-left:16px;padding-right:16px}._2pib{padding-left:20px;padding-right:20px}._2o1m{padding-left:24px;padding-right:24px}._2o1n{padding-left:36px;padding-right:36px}._iky{padding-top:0}._2pic{padding-top:4px}._2pid{padding-top:8px}._2pie{padding-top:12px}._2pif{padding-top:16px}._2pig{padding-top:20px}._2owm{padding-top:24px}._div{padding-right:0}._2pih{padding-right:4px}._2pii{padding-right:8px}._2pij{padding-right:12px}._2pik{padding-right:16px}._2pil{padding-right:20px}._31wk{padding-right:24px}._2phb{padding-right:32px}._au-{padding-bottom:0}._2pim{padding-bottom:4px}._2pin{padding-bottom:8px}._2pio{padding-bottom:12px}._2pip{padding-bottom:16px}._2piq{padding-bottom:20px}._2o1p{padding-bottom:24px}._4gao{padding-bottom:32px}._1cvx{padding-left:0}._2pir{padding-left:4px}._2pis{padding-left:8px}._2pit{padding-left:12px}._2piu{padding-left:16px}._2piv{padding-left:20px}._2o1q{padding-left:24px}._2o1r{padding-left:36px}.uiBoxGray{background-color:#f2f2f2;border:1px solid #ccc}.uiBoxDarkgray{color:#ccc;background-color:#333;border:1px solid #666}.uiBoxGreen{background-color:#d1e6b9;border:1px solid #629824}.uiBoxLightblue{background-color:#33d2fa;border:1px solid #d8dfea}.uiBoxRed{background-color:#ffebe8;border:1px solid #dd3c10}.uiBoxWhite{background-color:#fff;border:1px solid #ccc;margin-right:55%}.uiBoxBlue{background-color:#5627ff;border:1px solid #ccc;margin-left:55%}.uiBoxYellow{background-color:#fff9d7;border:1px solid #e2c822}.uiBoxOverlay{background:rgba(255,255,255,.85);border:1px solid #3b5998;border:1px solid rgba(59,89,153,.65);zoom:1}.noborder{border:0}.topborder{border-bottom:0;border-left:none;border-right:0}.bottomborder{border-left:none;border-right:0;border-top:0}.dashedborder{border-style:dashed}.pas{padding:5px}.pam{padding:10px;padding-left:20px;padding-right:20px;padding-top:10px}.pal{padding:20px}.pts{padding-top:5px}.ptm{padding-top:10px}.ptl{padding-top:20px}.prs{padding-right:5px}.prm{padding-right:10px}.prl{padding-right:20px}.pbs{padding-bottom:5px}.pbm{padding-bottom:10px}.pbl{padding-bottom:20px}.pls{padding-left:5px}.plm{padding-left:10px}.pll{padding-left:20px}.phs{padding-left:5px;padding-right:5px}.phm{padding-left:10px;padding-right:10px}.phl{padding-left:20px;padding-right:20px}.pvs{padding-top:5px;padding-bottom:5px}.pvm{padding-top:10px;padding-bottom:10px}.pvl{padding-top:20px;padding-bottom:20px}.mas{margin:5px}.mam{margin:10px}.mal{margin:20px}.mts{margin-top:5px}.mtm{margin-top:10px}.mtl{margin-top:20px}.mrs{margin-right:5px}.mrm{margin-right:10px}.mrl{margin-right:20px}.mbs{margin-bottom:5px}.mbm{margin-bottom:10px}.mbl{margin-bottom:20px}.mls{margin-left:5px}.mlm{margin-left:10px}.mll{margin-left:20px}.mhs{margin-left:5px;margin-right:5px}.mhm{margin-left:10px;margin-right:10px}.mhl{margin-left:20px;margin-right:20px}.mvs{margin-top:5px;margin-bottom:5px}.mvm{margin-top:10px;margin-bottom:10px}.mvl{margin-top:20px;margin-bottom:20px}._id9{float:left}._idm{float:right}._idn{float:none}._37no{font-size:13px;padding-bottom:8px}._u14{color:gray;font-size:13px;padding-bottom:8px}._12gz{font-size:14px;font-weight:bold;padding-bottom:8px}._67gx{color:gray;font-size:13px}._3bki{color:#90949c;font-size:13px}.Custom_Name{width:90%;margin-top:7px;margin-left:5px;color:#000;font-weight:bold;font-size:24px}p{margin:0;font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:23px;font-style:normal;font-variant:normal;font-weight:700;
		line-height:23px}h3{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:17px;font-style:normal;font-variant:normal;font-weight:700;line-height:23px}p{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:20px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}blockquote{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:17px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}pre{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:11px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}#buttonB{background-image:linear-gradient(to bottom right,#4c23e2,#04b8ff);border:0;color:white;padding:5px 20px;text-align:center;text-decoration:none;display:inline-block;font-size:16px;margin-left:1%;margin-right:.4%;margin-top:.4%;border-radius:10px;transition-duration:.4s;height:30px}#buttonB:hover{transform:scale(1.1)}img{height:100%;width:100%}</style>;`))

		//Handles any file write errors
		ErrHandle(wrF0err)

		//Creates the appropriate html page prefix
		prepTxt := fmt.Sprintf(`<title>%v</title> </head> <body class="_5vb_ _2yq _4yic"><div class="clearfix _ikh"> <div class="_4bl9"> <div class="_li"> <div id="bluebarRoot" class="_2t-8 _1s4v _2s1x _h2p _3b0a"> <div aria-label="Facebook" class="_2t-a _26aw _5rmj _50ti _2s1y" role="banner"> <p class="Custom_Name" id="namething">%v</p> </div> </div> </div> <div class="_3a_u"> <div class="_4t5n" role="main" id="message_box">`, name, name)

		//Writes the prefix to the file
		_, wrFerr := expFile.Write([]byte(prepTxt))

		//Handles any file write errors
		ErrHandle(wrFerr)

		for _, message := range jsonData[id].Conversation {

			timeOBJ, timeErr := time.Parse(time.RFC3339, message.Timestamp)

			ErrHandle(timeErr)

			timestmp := timeOBJ.Format("Mon Jan 2 15:04:05 2006")

			if message.Sender == master {

				switch {
				case message.MediaShareOwner != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div>Media: %v</div> <div>Caption: %v</div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaShareOwner, message.MediaShareCaption, message.MediaShareURL, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent1 != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent1, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.StoryShare != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div>%v</div> <div></div> </div> </div> <div class="_3-94 _2lem">%v</div> </div>`, getMSGtext(message), message.StoryShare, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				default:

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div> </div> </div> <div class="_3-94 _2lem">%s</div> </div>`, getMSGtext(message), timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				}

			} else {

				switch {
				case message.MediaShareOwner != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2let"> <div> <div></div> <div>%v</div> <div>Media: %v</div> <div>Caption: %v</div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaShareOwner, message.MediaShareCaption, message.MediaShareURL, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2let"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent1 != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2let"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent1, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.StoryShare != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2let"> <div> <div></div> <div>%v</div> <div>%v</div> <div></div> </div> </div> <div class="_3-94 _2lem">%v</div> </div>`, getMSGtext(message), message.StoryShare, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				default:

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2let"> <div> <div></div> <div>%v</div> <div></div> <div></div> </div> </div> <div class="_3-94 _2lem">%s</div> </div>`, getMSGtext(message), timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				}

			}

		}

	} else { //If the conversation is a group

		name = fmt.Sprintf("Group, %v participants", len(jsonData[id].Participants))
		groupFileName := fmt.Sprintf("%vgroup", id)

		if standalone {

			usedPath = &standMainPATH

		} else {

			usedPath = &mainPath

		}

		expFile := createFile(*usedPath, groupFileName)

		//Writes the prefix to the file
		_, wrF0err := expFile.Write([]byte(`<!DOCTYPE html>
		<html>
		<head>
		<meta http-equiv=Content-Type content="text/html; charset=UTF-8, user-scalable=yes" />
		<style type=text/css>body{background:#fff;color:#1c1e21;direction:ltr;line-height:1.34;margin:0;padding:0;unicode-bidi:embed}body,button,input,label,select,td,textarea{font-family:Helvetica,Arial,sans-serif;font-size:12px}h1,h2,h3,h4,h5,h6{color:#1c1e21;font-size:13px;font-weight:600;margin:0;padding:0}h1{font-size:14px}h4,h5,h6{font-size:12px}b,strong{font-weight:600}a{color:#385898;cursor:pointer;text-decoration:none}button{margin:0}a:hover{text-decoration:underline}img{border:0}td,td.label{text-align:left}dd{color:#000}dt{color:#606770}ul{list-style-type:none;margin:0;padding:0}abbr{border-bottom:0;text-decoration:none}hr{background:#dadde1;border-width:0;color:#dadde1;height:1px}form{margin:0;padding:0}label{color:#606770;cursor:default;font-weight:600;vertical-align:middle}label input{font-weight:normal}textarea,.inputtext,.inputpassword{border:1px solid #ccd0d5;border-radius:0;margin:0;padding:3px}textarea{max-width:100%}select{border:1px solid #ccd0d5;padding:2px}input,select,textarea{background-color:#fff;color:#1c1e21}.inputtext,.inputpassword{padding-bottom:4px}.inputtext:invalid,.inputpassword:invalid{box-shadow:none}.inputradio{margin:0 5px 0 0;padding:0;vertical-align:middle}.inputcheckbox{border:0;vertical-align:middle}.inputbutton,.inputsubmit{background-color:#4267b2;border-color:#dadde1 #0e1f5b #0e1f5b #d9dfea;border-style:solid;border-width:1px;color:#fff;padding:2px 15px 3px 15px;text-align:center}.inputsubmit_disabled{background-color:#999;border-bottom:1px solid #000;border-right:1px solid #666;color:#fff}.inputaux{background:#ebedf0;border-color:#ebedf0 #666 #666 #e7e7e7;color:#000}.inputaux_disabled{color:#999}.inputsearch{background:#fff url(https://static.xx.fbcdn.net/rsrc.php/v3/yV/r/IJYgcESal33.png) no-repeat left 4px;padding-left:17px}.clearfix:after{clear:both;content:'.';display:block;font-size:0;height:0;line-height:0;visibility:hidden}.clearfix{zoom:1}.datawrap{word-wrap:break-word}.word_break{display:inline-block}.flexchildwrap{min-width:0;word-wrap:break-word}.ellipsis{overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.aero{opacity:.5}.column{float:left}.center{margin-left:auto;margin-right:auto}#facebook .hidden_elem{display:none!important}#facebook .invisible_elem{visibility:hidden}#facebook .accessible_elem{clip:rect(1px,1px,1px,1px);height:1px;overflow:hidden;position:absolute;white-space:nowrap;width:1px}.direction_ltr{direction:ltr}.direction_rtl{direction:rtl}.text_align_ltr{text-align:left}.text_align_rtl{text-align:right}body{overflow-y:scroll}.mini_iframe,.serverfbml_iframe{overflow-y:visible}.auto_resize_iframe{height:auto;overflow:hidden}.pipe{color:gray;padding:0 3px}#content{margin:0;outline:0;padding:0;width:auto}.profile #content,.home #content,.search #content{min-height:600px}.UIStandardFrame_Container{margin:0 auto;padding-top:20px;width:960px}.UIStandardFrame_Content{float:left;margin:0;padding:0;width:760px}.UIStandardFrame_SidebarAds{float:right;margin:0;padding:0;width:200px;word-wrap:break-word}.UIFullPage_Container{margin:0 auto;padding:20px 12px 0;width:940px}.empty_message{background:#f5f6f7;font-size:13px;line-height:17px;padding:20px 20px 50px;text-align:center}.see_all{text-align:right}.standard_status_element{visibility:hidden}.standard_status_element.async_saving{visibility:visible}img.tracking_pixel{height:1px;position:absolute;visibility:hidden;width:1px}#globalContainer{margin:0 auto;position:relative;zoom:1}.fbx #globalContainer{width:981px}.sidebarMode #globalContainer{padding-right:205px}.fbx #tab_canvas>div{padding-top:0}.fb_content{min-height:612px;padding-bottom:20px}.fbx .fb_content{padding-bottom:0}.skipto{display:none}.home .skipto{display:block}._li._li._li{overflow:initial}._72b0{position:relative;z-index:0}._5vb_ #pageFooter{display:none}html body._5vb_ #globalContainer{width:976px}._5vb_.hasLeftCol #headerArea{margin:0;padding-top:0;width:786px}._5vb_,._5vb_ #contentCol{background-color:#e9ebee;color:#1d2129}html ._5vb_.hasLeftCol #contentCol{border-left:0;margin-left:172px;padding-left:11px;padding-top:11px}._5vb_.hasLeftCol #topNav{border-left:0;margin-left:172px;padding:11px 7px 0 11px}._5vb_.hasLeftCol #topNav~#contentCol{padding-top:0}._5vb_.hasLeftCol #leftCol{padding-left:8px;padding-top:12px;width:164px}._5vb_.hasLeftCol #mainContainer{border-right:0;margin-left:0}._5vb_.hasLeftCol #pageFooter{background:0}html ._5vb_._5vb_.hasLeftCol div#contentArea{padding-left:0;padding-right:10px;width:786px}html ._5vb_._5vb_.hasLeftCol .hasRightCol div#contentArea{width:496px}._5vb_.hasLeftCol ._5r-_ div#rightCol{padding:0 7px 0 0;width:280px}._2yq #globalContainer{width:1012px!important}._2yq #headerArea{float:none!important;padding:0 0 12px 0!important;width:auto!important}._2yq #contentArea{margin-right:0;padding:0!important}._2yq #leftCol,._2yq #contentCol{padding:0!important}._2yq #rightCol{float:left;margin-top:0;padding:0!important}._2yq .groupJumpLayout{margin-top:-12px}._2yq .loggedout_menubar_container{min-width:1014px}._4yic{margin:auto;font-size:13px}._3a_u{margin:0 auto;width:50%}._4t5n{float:none;margin-bottom:15px;position:relative;word-break:break-word;z-index:0}._4t5o{clear:both;color:#7f7f7f;font-size:14px;margin-bottom:20px;margin-top:10px;text-align:center}._3b0a{position:relative;z-index:100}._3b0b{background:#fff;border-radius:3px;display:flex;flex-direction:row;padding:15px}._3z-t{border-radius:50%;height:16px;padding:4px;width:16px}._3b0c{display:flex;flex-direction:column;justify-content:center;margin-left:8px}._3b0d{color:#1d2129;font-size:14px;font-weight:bold;line-height:18px;margin-bottom:3px}._3b0e{color:#90949c;font-size:14px;line-height:16px}._218o{display:flex;justify-content:space-between;margin:0 auto;margin-left:60%;width:auto}._72m4{font-size:12px;margin-top:4px}._5aj7{display:flex}._5aj7 ._4bl7{float:none}._5aj7 ._4bl9{flex:1 0 0}._ikh ._4bl7{float:left;min-height:1px}._4bl7,._4bl9{word-wrap:break-word}._4bl9{overflow:hidden}._21dp{position:relative;z-index:301}._2t-8._2t-8{font-family:Helvetica,Arial,sans-serif}._2t-8{height:43px;min-width:100%}._2t-a{height:42px;position:relative;width:100%}._2s1y{box-sizing:border-box;height:43px}._50ti{position:fixed;top:0}.hasDemoBar ._50ti{top:60px}._33rf{min-width:981px}._2yq ._33rf,._2xk0 ._33rf{min-width:1014px}._50tj{box-sizing:border-box;padding-right:0}.sidebarMode ._50tj{padding-right:205px}._4pmj{box-sizing:border-box;padding:0 16px}._2t-d{margin:auto;padding:0 8px}._2t-a{display:flex;justify-content:space-between}._2t-e,._2t-f{display:flex}._2t-e{flex:1 1 auto;justify-content:flex-start}._2t-f{flex:0 0 auto;justify-content:flex-end;margin-left:8px}._4kny{float:left}._50tm{width:100%}._2s24{margin-left:1px;position:relative}._h2p ._2s24{margin-left:0}._cy6{display:inline-block;padding:0 9px;vertical-align:top}._h2p ._cy6{padding:0 5px 0 4px}._cy6:first-child{padding-left:0}._h2p ._cy6:first-child{padding:0}._cy6:last-child{padding-right:0}._cy7{margin:7px 0 8px 0}._h2p ._cy7{margin-right:1px}._2s25{background-color:transparent;color:inherit;display:inline-block;font-size:12px;font-weight:bold;height:27px;line-height:28px;padding:0 10px 1px;position:relative;text-decoration:none;vertical-align:top;white-space:nowrap}.segoe ._2s25{font-weight:600}._h2p ._2s25{padding:0 12px 1px}._h2p ._cy6 ._4kny:last-child ._2s25{padding-right:11px}._2s25:hover,._2s25:focus,._2s25:active{border-radius:2px;color:inherit;outline:0;text-decoration:none;z-index:1}.openToggler ._2s25:hover,.openToggler ._2s25:focus,.openToggler ._2s25:active{background:transparent}._4kny ._585-{margin-left:0;min-width:144px;width:100%}._4kny .__wu ._539-.roundedBox{margin-left:0}._4kny ._4962{float:none;margin:5px 0 6px 0;position:relative}._h2p ._4kny ._4962{margin:5px 0 6px}._3x1p{height:100%;overflow:hidden;position:absolute;width:100%}._63i8{display:block;left:0;opacity:0;position:absolute;top:0;transition:opacity 200ms;width:100vw;z-index:2}._4yim,._4yin{display:block;left:0;position:absolute;top:0;transform-origin:left}._4yin{transform-origin:right}._4yio,._4yip{display:inline-block;height:2px;left:0;min-width:12px;position:absolute;top:0;transform-origin:left}._4yip{transform-origin:right}._2t-8.indeterminateBarTransition.transitioning ._63i8{opacity:1}._2t-8.indeterminateBarTransition.transitioning ._4yim{animation:indeterminateBarTransitionTranslate-left 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yin{animation:indeterminateBarTransitionTranslate-right 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yio{animation:indeterminateBarTransitionWidth-left 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.transitioning ._4yip{animation:indeterminateBarTransitionWidth-right 2000ms infinite;animation-timing-function:steps(20,end)}._2t-8.indeterminateBarTransition.finishing ._63i8{opacity:0}@keyframes indeterminateBarTransitionTranslate-left{0%{animation-timing-function:ease-in;transform:translateX(0)}25%{animation-timing-function:ease-out;transform:translateX(25vw)}50%{animation-timing-function:ease-in;transform:translateX(calc(100vw - 12px))}100%{transform:translateX(calc(100vw - 12px))}}@keyframes indeterminateBarTransitionTranslate-right{0%{transform:translateX(0)}51%{animation-timing-function:ease-in;transform:translateX(0)}75%{animation-timing-function:ease-out;transform:translateX(calc(-25vw))}100%{animation-timing-function:ease-in;transform:translateX(calc(-100vw+12px))}}@keyframes indeterminateBarTransitionWidth-left{0%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}25%{animation-timing-function:ease-out;transform:scaleX(50)}50%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}51%{opacity:0}100%{opacity:0}}@keyframes indeterminateBarTransitionWidth-right{0%{opacity:0}25%{opacity:0}50%
		{animation-timing-function:ease-in;opacity:0;transform:scaleX(1)}51%{opacity:1}75%{animation-timing-function:ease-out;opacity:1;transform:scaleX(50)}100%{animation-timing-function:ease-in;opacity:1;transform:scaleX(1)}}._2s1x ._2s1y{background-color:#fff;color:#fff;box-shadow:0 4px 5px 0 rgba(179,179,179,1)}._2s1x ._2s24::before{background:rgba(0,0,0,.1)}._2s1x ._2s25:hover,._2s1x ._2s25:focus,._2s1x ._2s25:active{background:rgba(0,0,0,.1);color:inherit}._2s1x.transitioning ._3fju{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yv/r/z1PAxf53vph.gif);height:100%;width:100%}._2s1x.transitioning ._3fjv{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/y5/r/OzkCShPcfVN.gif);height:100%;width:100%}._2s1x.transitioning ._3fjx{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yH/r/YIs0iw6cHAa.gif);height:100%;width:100%}._2s1x ._3b33{background:#fff;height:100%;opacity:0;position:absolute;width:100%}._2s1x.pulseTransition.transitioning ._3b33{animation:pulse-loading 800ms cubic-bezier(.455,.03,.515,.955) infinite alternate;animation-timing-function:steps(8,end)}@keyframes pulse-loading{0%{opacity:0}100%{opacity:.15}}._2s1x ._3b34{height:100%;max-width:1016px;position:absolute;width:100%}._2s1x.shimmerTransition.transitioning ._3b34{animation:shimmer-loading 1600ms cubic-bezier(.455,.03,.515,.955) infinite;animation-timing-function:steps(16,end);background:linear-gradient(to right,rgba(66,103,178,0),#577fbc,rgba(66,103,178,0));background-size:1016px auto}@keyframes shimmer-loading{0%{transform:translateX(-1016px)}100%{transform:translateX(calc(100vw+1016px))}}._2s1x ._63tk{background-color:#fff}._19ea{margin:7px 0;margin-left:-2px;margin-right:5px}._19eb{display:inline-block;outline:0;padding:2px;position:relative}._7tp1{display:inline-block;height:20px;outline:0;padding:4px;position:relative}._19eb:hover,._19eb:focus,._19eb:active,._7tp1:hover,._7tp1:focus,._7tp1:active{background-color:#365899;background-color:rgba(0,0,0,.1);border-radius:3px}._2md{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yr/r/Xo5dkI2ODGj.png);background-repeat:no-repeat;background-size:auto;background-position:-25px -116px;display:block;height:24px;outline:0;overflow:hidden;text-indent:-999px;white-space:nowrap;width:24px}._7tp2{background-image:url(https://static.xx.fbcdn.net/rsrc.php/v3/yr/r/Xo5dkI2ODGj.png);background-repeat:no-repeat;background-size:auto;background-position:0 0;display:block}._7ql{border-radius:2px;display:inline;margin:2px 6px 2px -8px;vertical-align:inherit}._h2p ._7ql{margin-left:-10px}._1k67 ._2s25{position:relative}._1k67._d0b ._2s25{padding-right:0}._1k67 ._1vp5 .img{transform:translateY(2px)}._1k67._d0b._5-y2 ._2s25{padding-right:6px}._2qgu._2qgu{border-radius:50%;overflow:hidden}._2s25._2s25._606w._606w:after,._606w:after{border-radius:50%}._605a .fbxWelcomeBoxBlock:after{border-radius:50%}._1qv9{align-items:center;display:flex;flex-direction:row}._rv{height:100px;width:100px}._rw{height:50px;width:50px}._s0:only-child{display:block}._3tm9{height:14px;width:14px}._54rv{height:16px;width:16px}._3qxe{height:19px;width:19px}._1m6h{height:24px;width:24px}._3d80{height:28px;width:28px}._54ru{height:32px;width:32px}._tzw{height:40px;width:40px}._54rt{height:48px;width:48px}._54rs{height:56px;width:56px}._1m9m{height:64px;width:64px}._ry{height:24px;width:24px}._4jnw{margin:0}._3-8h{margin:4px}._3-8i{margin:8px}._3-8j{margin:12px}._3-8k{margin:16px}._3-8l{margin:20px}._2-5b{margin:24px}._1kbd{margin-bottom:0;margin-top:0}._3-8m{margin-bottom:4px;margin-top:4px}._3-8n{margin-bottom:8px;margin-top:8px}._3-8o{margin-bottom:12px;margin-top:12px}._3-8p{margin-bottom:16px;margin-top:16px}._3-8q{margin-bottom:20px;margin-top:20px}._2-ox{margin-bottom:24px;margin-top:24px}._1a4i{margin-left:0;margin-right:0}._3-8r{margin-left:4px;margin-right:4px}._3-8s{margin-left:8px;margin-right:8px}._3-8t{margin-left:12px;margin-right:12px}._3-8u{margin-left:16px;margin-right:16px}._3-8v{margin-left:20px;margin-right:20px}._6bu9{margin-left:24px;margin-right:24px}._5soe{margin-top:0}._3-8w{margin-top:4px}._3-8x{margin-top:8px}._3-8y{margin-top:12px}._3-8z{margin-top:16px}._3-8-{margin-top:20px}._4aws{margin-top:24px}._2-jz{margin-right:0}._3-8_{margin-right:4px}._3-90{margin-right:8px}._3-91{margin-right:12px}._3-92{margin-right:16px}._3-93{margin-right:20px}._y8t{margin-right:24px}._5emk{margin-bottom:0}._3-94{margin-bottom:4px}._3-95{margin-bottom:8px}._3-96{margin-bottom:12px}._3-97{margin-bottom:16px}._3-98{margin-bottom:20px}._20nr{margin-bottom:24px}._av_{margin-left:0}._3-99{margin-left:4px}._3-9a{margin-left:8px}._3-9b{margin-left:12px}._3-9c{margin-left:16px}._3-9d{margin-left:20px}._4m0t{margin-left:24px}._2lej{border-radius:20px}._2lek{color:#1d2129;font-size:14px;font-weight:bold;line-height:18px}._2lak{color:#fff;font-size:14px;font-weight:bold;line-height:18px}._2lel{border-bottom:1px solid #dadde1}._2lem,._2lem a{color:#8d949e;font-size:13px;line-height:16px}._2let{color:#1d2129;font-size:18px;line-height:22px}._2lat{color:#fff;font-size:18px;line-height:22px}._tqp{color:gray;font-size:13px}._4mp8{font-weight:bold}._4nkx,._3ttj{font-size:14px;line-height:2;text-align:left}._4nkx tbody tr th{padding:5px 5px;text-align:left;vertical-align:top;width:150px}._2yuc{max-width:100%}._3hls{font-size:14px;font-weight:bold}._2oao{color:#90949c;font-size:13px;font-weight:bold;line-height:20px;width:100px}._23bw{font-size:13px}._6udd{word-break:break-all}._8tm{padding:0}._2phz{padding:4px}._2ph-{padding:8px}._2ph_{padding:12px}._2pi0{padding:16px}._2pi1{padding:20px}._40c7{padding:24px}._2o1j{padding:36px}._6buq{padding-bottom:0;padding-top:0}._2pi2{padding-bottom:4px;padding-top:4px}._2pi3{padding-bottom:8px;padding-top:8px}._2pi4{padding-bottom:12px;padding-top:12px}._2pi5{padding-bottom:16px;padding-top:16px}._2pi6{padding-bottom:20px;padding-top:20px}._2o1k{padding-bottom:24px;padding-top:24px}._2o1l{padding-bottom:36px;padding-top:36px}._6bua{padding-left:0;padding-right:0}._2pi7{padding-left:4px;padding-right:4px}._2pi8{padding-left:8px;padding-right:8px}._2pi9{padding-left:12px;padding-right:12px}._2pia{padding-left:16px;padding-right:16px}._2pib{padding-left:20px;padding-right:20px}._2o1m{padding-left:24px;padding-right:24px}._2o1n{padding-left:36px;padding-right:36px}._iky{padding-top:0}._2pic{padding-top:4px}._2pid{padding-top:8px}._2pie{padding-top:12px}._2pif{padding-top:16px}._2pig{padding-top:20px}._2owm{padding-top:24px}._div{padding-right:0}._2pih{padding-right:4px}._2pii{padding-right:8px}._2pij{padding-right:12px}._2pik{padding-right:16px}._2pil{padding-right:20px}._31wk{padding-right:24px}._2phb{padding-right:32px}._au-{padding-bottom:0}._2pim{padding-bottom:4px}._2pin{padding-bottom:8px}._2pio{padding-bottom:12px}._2pip{padding-bottom:16px}._2piq{padding-bottom:20px}._2o1p{padding-bottom:24px}._4gao{padding-bottom:32px}._1cvx{padding-left:0}._2pir{padding-left:4px}._2pis{padding-left:8px}._2pit{padding-left:12px}._2piu{padding-left:16px}._2piv{padding-left:20px}._2o1q{padding-left:24px}._2o1r{padding-left:36px}.uiBoxGray{background-color:#f2f2f2;border:1px solid #ccc}.uiBoxDarkgray{color:#ccc;background-color:#333;border:1px solid #666}.uiBoxGreen{background-color:#d1e6b9;border:1px solid #629824}.uiBoxLightblue{background-color:#33d2fa;border:1px solid #d8dfea}.uiBoxRed{background-color:#ffebe8;border:1px solid #dd3c10}.uiBoxWhite{background-color:#fff;border:1px solid #ccc;margin-right:55%}.uiBoxBlue{background-color:#5627ff;border:1px solid #ccc;margin-left:55%}.uiBoxYellow{background-color:#fff9d7;border:1px solid #e2c822}.uiBoxOverlay{background:rgba(255,255,255,.85);border:1px solid #3b5998;border:1px solid rgba(59,89,153,.65);zoom:1}.noborder{border:0}.topborder{border-bottom:0;border-left:none;border-right:0}.bottomborder{border-left:none;border-right:0;border-top:0}.dashedborder{border-style:dashed}.pas{padding:5px}.pam{padding:10px;padding-left:20px;padding-right:20px;padding-top:10px}.pal{padding:20px}.pts{padding-top:5px}.ptm{padding-top:10px}.ptl{padding-top:20px}.prs{padding-right:5px}.prm{padding-right:10px}.prl{padding-right:20px}.pbs{padding-bottom:5px}.pbm{padding-bottom:10px}.pbl{padding-bottom:20px}.pls{padding-left:5px}.plm{padding-left:10px}.pll{padding-left:20px}.phs{padding-left:5px;padding-right:5px}.phm{padding-left:10px;padding-right:10px}.phl{padding-left:20px;padding-right:20px}.pvs{padding-top:5px;padding-bottom:5px}.pvm{padding-top:10px;padding-bottom:10px}.pvl{padding-top:20px;padding-bottom:20px}.mas{margin:5px}.mam{margin:10px}.mal{margin:20px}.mts{margin-top:5px}.mtm{margin-top:10px}.mtl{margin-top:20px}.mrs{margin-right:5px}.mrm{margin-right:10px}.mrl{margin-right:20px}.mbs{margin-bottom:5px}.mbm{margin-bottom:10px}.mbl{margin-bottom:20px}.mls{margin-left:5px}.mlm{margin-left:10px}.mll{margin-left:20px}.mhs{margin-left:5px;margin-right:5px}.mhm{margin-left:10px;margin-right:10px}.mhl{margin-left:20px;margin-right:20px}.mvs{margin-top:5px;margin-bottom:5px}.mvm{margin-top:10px;margin-bottom:10px}.mvl{margin-top:20px;margin-bottom:20px}._id9{float:left}._idm{float:right}._idn{float:none}._37no{font-size:13px;padding-bottom:8px}._u14{color:gray;font-size:13px;padding-bottom:8px}._12gz{font-size:14px;font-weight:bold;padding-bottom:8px}._67gx{color:gray;font-size:13px}._3bki{color:#90949c;font-size:13px}.Custom_Name{width:90%;margin-top:7px;margin-left:5px;color:#000;font-weight:bold;font-size:24px}p{margin:0;font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:23px;font-style:normal;font-variant:normal;font-weight:700;
		line-height:23px}h3{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:17px;font-style:normal;font-variant:normal;font-weight:700;line-height:23px}p{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:20px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}blockquote{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:17px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}pre{font-family:"Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;font-size:11px;font-style:normal;font-variant:normal;font-weight:400;line-height:23px}#buttonB{background-image:linear-gradient(to bottom right,#4c23e2,#04b8ff);border:0;color:white;padding:5px 20px;text-align:center;text-decoration:none;display:inline-block;font-size:16px;margin-left:1%;margin-right:.4%;margin-top:.4%;border-radius:10px;transition-duration:.4s;height:30px}#buttonB:hover{transform:scale(1.1)}img{height:100%;width:100%}</style>;`))

		//Handles any file write errors
		ErrHandle(wrF0err)

		//Creates the appropriate html page prefix
		prepTxt := fmt.Sprintf(`<title>%v</title> </head> <body class="_5vb_ _2yq _4yic"><div class="clearfix _ikh"> <div class="_4bl9"> <div class="_li"> <div id="bluebarRoot" class="_2t-8 _1s4v _2s1x _h2p _3b0a"> <div aria-label="Facebook" class="_2t-a _26aw _5rmj _50ti _2s1y" role="banner"> <p class="Custom_Name" id="namething">%v</p> </div> </div> </div> <div class="_3a_u"> <div class="_4t5n" role="main" id="message_box">`, name, name)

		//Writes the prefix to the file
		_, wrFerr := expFile.Write([]byte(prepTxt))

		//Handles any file write errors
		ErrHandle(wrFerr)

		for _, message := range jsonData[id].Conversation {

			timeOBJ, timeErr := time.Parse(time.RFC3339, message.Timestamp)

			ErrHandle(timeErr)

			timestmp := timeOBJ.Format("Mon Jan 2 15:04:05 2006")

			if message.Sender == master {

				switch {
				case message.MediaShareOwner != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div>Media: %v</div> <div>Caption: %v</div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaShareOwner, message.MediaShareCaption, message.MediaShareURL, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent1 != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, getMSGtext(message), message.MediaSent1, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.StoryShare != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div>%v</div> <div></div> </div> </div> <div class="_3-94 _2lem">%v</div> </div>`, getMSGtext(message), message.StoryShare, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				default:

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div> </div> </div> <div class="_3-94 _2lem">%s</div> </div>`, getMSGtext(message), timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				}

			} else {

				switch {
				case message.MediaShareOwner != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2lek">%v</div><div class="_3-96 _2let"> <div></div> <div>%v</div> <div>Media: %v</div> <div>Caption: %v</div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, message.Sender, getMSGtext(message), message.MediaShareOwner, message.MediaShareCaption, message.MediaShareURL, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2lek">%v</div><div class="_3-96 _2let"> <div></div> <div>%v</div> <div></div> <div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, message.Sender, getMSGtext(message), message.MediaSent, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.MediaSent1 != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2lek">%v</div><div class="_3-96 _2let"> <div></div> <div>%v</div> <div></div> <div></div> </div> <div class="_3-94 _2lem"><img src="%v" alt="Image|Video not available"><br>%v</div> </div>`, message.Sender, getMSGtext(message), message.MediaSent1, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				case message.StoryShare != "":

					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2lek">%v</div><div class="_3-96 _2let"> <div></div> <div>%v</div> <div>%v</div> <div></div>  </div> <div class="_3-94 _2lem">%v</div> </div>`, message.Sender, getMSGtext(message), message.StoryShare, timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				default:
					/*<div class="pam _3-95 _2pi0 _2lej uiBoxBlue noborder"> <div class="_3-96 _2lat"> <div> <div></div> <div>%v</div> <div></div> <div></div> </div> </div> <div class="_3-94 _2lem">%s</div> </div>*/
					writeSRT := fmt.Sprintf(`<div class="pam _3-95 _2pi0 _2lej uiBoxWhite noborder"> <div class="_3-96 _2lek">%v</div><div class="_3-96 _2let"> <div></div> <div>%v</div> <div></div> <div></div> </div> <div class="_3-94 _2lem">%s</div> </div>`, message.Sender, getMSGtext(message), timestmp)

					_, wrMesERR := expFile.Write([]byte(writeSRT))

					//Handles any file write errors
					ErrHandle(wrMesERR)

				}

			}

		}

		endTxt := `</div> </div> <div class="_4t5o"> <div>End Of Documment</div> </aiv> </div> </div> </div> </body> </html>`

		//Writes the prefix to the file
		_, wrENDFerr := expFile.Write([]byte(endTxt))

		//Handles any file write errors
		ErrHandle(wrENDFerr)

	}

}

func exportAll() {
	var isGroup bool
	var name string

	allPath := createDir(false)

	pathByte := []byte(allPath)
	newPathB := pathByte[:len(pathByte)-8]

	newPath := string(newPathB)

	for i := range jsonData {
		fmt.Printf("Exporting %v...", i)
		exportConv(i, false, allPath)
		fmt.Printf(" Done!\n")
	}

	indexFile := createFile(newPath, "index")

	_, idxwrErr := indexFile.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8, user-scalable=yes" /><title>Messages</title><style>#buttonBox {width: 75%;margin-top: 20%;margin-left: 12.5%;margin-right: 12.5%;background-color: transparent;}#buttonCon {font-family: Arial; background-image: linear-gradient(to bottom right,#353535, #868686);border: none;color: white;padding: 20px 1%;text-align: center;text-decoration: none;display: inline-block;font-size: 20px;margin-left: 1%;margin-right: 1%;margin-top: 2%;border-radius: 16px;transition-duration: 0.4s}#buttonCon:hover {transform: scale(1.1)}#buttonGroup {font-family: Arial; background-image: linear-gradient(to bottom right, rgb(255, 0, 119), rgb(255, 166, 0));border: none;color: white;padding: 20px 1%;text-align: center;text-decoration: none;display: inline-block;font-size: 20px;margin-left: 1%;margin-right: 1%;margin-top: 2%;border-radius: 16px;transition-duration: 0.4s}#buttonGroup:hover {transform: scale(1.1)}#printable {margin-left: 2%;margin-right: 2%}</style></head><body id="body"><div style="text-align: center"><div style="display: inline-block" id="printable">`)

	ErrHandle(idxwrErr)

	for i, block := range jsonData {

		for _, participant := range block.Participants {
			if participant == master {
				continue
			}
			if len(block.Participants) >= 3 {
				isGroup = true
			}
			name = participant
		}
		if !isGroup {
			writeSRT := fmt.Sprintf(`<a id="buttonCon" href="messages/%v%v.html">%v</a>`, i, name, name)

			_, err527 := indexFile.WriteString(writeSRT)

			ErrHandle(err527)
		} else {
			writeSRT := fmt.Sprintf(`<a id="buttonGroup" href="messages/%vgroup.html">%v members</a>`, i, len(block.Participants))

			_, err529 := indexFile.WriteString(writeSRT)

			ErrHandle(err529)

			isGroup = false
		}
	}

	_, err541 := indexFile.WriteString("</div></div></body></html>")

	ErrHandle(err541)
}

func main() {
	endProg := false

	fmt.Println("Insta DM json parser v1.0")

	openf()

	for !endProg {
		var icmd string

		fmt.Println(`Enter cmd: (help, list, export, fetch, closef, closep)`)
		fmt.Scanln(&icmd)

		switch icmd {
		case "help":
			//---//
			fmt.Println("")
			fmt.Println("Showing help: [cmd] - function")
			fmt.Println("[list] - lists all the participants in the file with their id ([id] - @tag)")
			fmt.Println("[export] - exports all conversations")
			fmt.Println("[fetch] - exports a single convarsation using it's id")
			fmt.Println("[closef] - closes the current file and promts you to open another one")
			fmt.Println("[closep] - closes the current file and terminates the program")
			fmt.Println("")
			//---//
		case "list":
			//---//
			list()
			//---//
		case "export":

			fmt.Println("")
			exportAll()
			fmt.Println("")
			fmt.Println("Done!!")
			fmt.Println("")

		case "fetch":
			var inID int
			fmt.Scanln(&inID)
			exportConv(inID, true, "C:\\MaPath\\files\\")

		case "closef":
			//---//
			fmt.Println("")
			openf()
			//---//
		case "closep":
			//---//
			fmt.Println("")
			fmt.Println("Closing program...")
			endProg = true
			//---//
		default:
			//---//
			fmt.Println("")
			fmt.Printf("Unknown cmd %q\n\n", icmd)
			//---//
		}

	}

	/*       //Struct Parse Test

	file, fcreateErr := os.Create("outdata.txt")

	ErrHandle(fcreateErr)

	defer file.Close()


	fmt.Println("Writing to File....")

	str := fmt.Sprintf("%+v", data)

	wrbytes, wrErr := file.WriteString(str)

	ErrHandle(wrErr)

	fmt.Println("Wrote %d bytes", wrbytes)

	*/
}