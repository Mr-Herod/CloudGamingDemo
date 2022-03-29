/* eslint-env browser */
function getCookie(cname)
{
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++)
    {
        var c = ca[i].trim();
        if (c.indexOf(name)==0) return c.substring(name.length,c.length);
    }
    return "";
}

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
    gameWin(e.data)
}

function gameWin(sc){
	alert("你真棒！本次得分: "+sc)
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
    el.style.height = "720px"
    el.style.width = "1536px"
    el.style.position = "fixed"
    el.style.zIndex = "0"
    el.style.left = "0"
    el.style.top = "0"
    let t1=document.getElementById('loadingText');//选取id为test的元素
    t1.style.display = 'none';	// 隐藏选择的元素
    let t2=document.getElementById('loadingImg');//选取id为test的元素
    t2.style.display = 'none';	// 隐藏选择的元素
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
	dataJSON["clientDes"] = des;
    dataJSON["username"] = getCookie("username");
    dataJSON["nickname"] = getCookie("nickname");
    dataJSON["gamename"] = "推箱子";

	$.ajax({
		url:"/startGame",
		type:"POST",
        withCredential:true,
		contentType: "application/json;charset=utf-8",
		data:JSON.stringify(dataJSON),
		success:function(resData){
            // debugger
			// alert("操作成功");
            // log("data:"+resData);
			setDes(resData);
		},
		error:function(XMLHttpRequest, textStatus, errorThrown){
            // debugger
            console.log(XMLHttpRequest.status);
            console.log(XMLHttpRequest.readyState);
		    console.log(textStatus);
			alert("操作失败 请稍后重试！");
            window.location.href= "/center"
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
