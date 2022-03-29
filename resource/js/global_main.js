function getCookie(cname)
{
    let name = cname + "=";
    let ca = document.cookie.split("; ");
    console.log( document.cookie)
    for(let i=0; i<ca.length; i++)
    {
        let c = ca[i].trim();
        if (c.indexOf(name)==0)
            return c.substring(name.length,c.length);
    }
    return "";
}

function welcome(){
    document.getElementById("welcomeBox").innerText="你好，"+getCookie("nickname")+";"
}