function CreateLine(container,options){
    const computedStyle = getComputedStyle(container)

    const width = computedStyle.width
    const height = computedStyle.fontSize
    const intwidth = parseInt(width)
    const intheight = parseInt(height)

    var svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("width", "100%");
    svg.setAttribute("height", "1em");

    const rc = rough.svg(svg)
    let node = rc.line(0,intheight/2,intwidth,intheight/2,options)
    svg.append(node)
    return svg
}

function CreateRect(container,options,withoutBorder){
    const computedStyle = getComputedStyle(container)

    const width = computedStyle.width
    const height = computedStyle.height
    const fontSize = computedStyle.fontSize
    const intwidth = parseInt(width)
    const intheight = parseInt(height)
    const intfontSize = parseInt(fontSize)

    var svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("width", intwidth + intfontSize);
    svg.setAttribute("height", intheight + intfontSize);


    const rc = rough.svg(svg)
    // let node = rc.line(0,intheight/2,intwidth,intheight/2,options)
    let node = rc.rectangle(intfontSize/2, intfontSize/2, intwidth,intheight,options)

    if(withoutBorder){


        node.removeChild(node.lastElementChild)
    }

    svg.append(node)
    return svg
}


let options ={
    strokeWidth:1,
    roughness:1,
    bowing:1.5
}
let containers = document.querySelectorAll(".rough-hr-container")
containers.forEach(container => {
    container.append(CreateLine(container,options))
});

options ={
        strokeWidth:0.3,
        strokeLineDash:[5, 5],
}

let navContainer = document.querySelector("article nav")
navContainer.append(CreateRect(navContainer,options))


options={
    fillStyle:"solid",
    fill:"#ebdbb2",
    roughness:3,
    bowing:3,
}
let backhomeContainer = document.querySelector(".back-to-home")
backhomeContainer.append(CreateRect(backhomeContainer,options,true))

let toc = document.querySelector("article nav")
let sidetocContainer = document.querySelector(".side-toc-container")
sidetocContainer.append(toc.cloneNode(true))

let floattocButton = document.querySelector(".float-toc")

floattocButton.onclick = ()=>{
    console.log(sidetocContainer.style.visibility);
    if( sidetocContainer.style.visibility == "" ||  sidetocContainer.style.visibility == "visible"){
        console.log("test visible succe");
    sidetocContainer.style.visibility = "hidden"
    return
    }
    if(sidetocContainer.style.visibility == "hidden"){
        sidetocContainer.style.visibility = "visible"
        return
    }
}