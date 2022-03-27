/* eslint-env browser */

const pc = new RTCPeerConnection({
  iceServers: [{
    urls: 'stun:stun.l.google.com:19302'
  }]
})


const log = msg => {
  document.getElementById('div').innerHTML += msg + '<br>'
}

const sendChannel = pc.createDataChannel('cgData')
sendChannel.onclose = () => console.log('sendChannel has closed')
sendChannel.onopen = () => console.log('sendChannel has opened')
sendChannel.onmessage = e => receiveMessage(e)

function receiveMessage(e){
	if (e.data == 'win') {
		gameWin()
	} else {
		log(`Message from DataChannel '${sendChannel.label}' payload '{${e.data}}'`)
	}
}

function gameWin(){
	alert("Cool you win!")
   	sendChannel.send(' ')
}

window.onkeypress=function(e){
   key = String.fromCharCode(e.keyCode)   
   sendChannel.send(key)
}


pc.ontrack = function (event) {
  const el = document.createElement(event.track.kind)
  el.srcObject = event.streams[0]
  el.autoplay = true
  el.loop = true
  //el.controls = true
  el.id = "videoWindow"
  el.muted = "muted"
  el.style.height = 80+"vh"
  el.style.witdth = 80+"vw"
  document.getElementById('remoteVideos').appendChild(el)
  el.play()
}

function setDes(des){
  const sd = des
  if (sd === '') {
    return alert('Session Description must not be empty')
  }

  try {
    // document.getElementById('remoteSessionDescription').value = des
    pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
  } catch (e) {
    alert(e)
  }
}

function sendDes(des){
	var dataJSON = {};
	dataJSON["des"] = des;
	$.ajax({
		url:"/sendDes",
		type:"POST",
		contentType: "application/json;charset=utf-8",
		data:JSON.stringify(dataJSON),
		success:function(resData){
           		// debugger
			// alert("操作成功");
                        // log("data:"+resData);
			setDes(resData);
		},
		error:function(XMLHttpRequest, textStatus, errorThrown){
            debugger
            // 状态码
            console.log(XMLHttpRequest.status);
            // 状态
            console.log(XMLHttpRequest.readyState);
            // 错误信息
		console.log(textStatus);
			alert("操作失败 请稍后重试！");
		},
	})	
}


pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
pc.onicecandidate = event => {
  if (event.candidate === null) {
    des = btoa(JSON.stringify(pc.localDescription))
    //document.getElementById('localSessionDescription').value = des
    sendDes(des) 
  }
}

// Offer to receive 1 audio, and 1 video track
pc.addTransceiver('video', {
  direction: 'sendrecv'
})


pc.createOffer().then(d => pc.setLocalDescription(d)).catch(log)
