
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
let containers = document.querySelectorAll(".wrapper .rough-hr-container")
containers.forEach(container => {
    container.append(CreateLine(container,options))
});

let pagenumbersContainers = document.querySelectorAll(".page-number")

options={
        strokeWidth:1,
        fill:"#fbf1c7",
        fillStyle:"solid"
}
pagenumbersContainers.forEach(pagenumbersContainer=>{
    pagenumbersContainer.append(CreateRect(pagenumbersContainer,options))
})

 options ={
        strokeWidth:1,
        fill:"black",
        fillStyle:"solid"
    }
pagenumbersContainers.forEach(pagenumbersContainer=>{
    pagenumbersContainer.append(CreateRect(pagenumbersContainer,options))
}
)

 options ={
        strokeWidth:0.8,
        fill:"#bdae93",
        fillStyle:"cross-hatch"
}

let tagsLiContainer = document.querySelectorAll(".tags-container li")

let withoutBorder = true
tagsLiContainer.forEach(li=>{
    li.append(CreateRect(li,options,withoutBorder))
})

options = {
    strokeWidth:1
}

let categoryContainer = document.querySelector(".category-container")
categoryContainer.append(CreateRect(categoryContainer,options))

let categoryLiContainer = document.querySelectorAll(".category-container li")
categoryLiContainer.forEach(li=>{
    li.append(CreateLine(li))
})

options = {
    strokeWidth:1
}

let tagsContainer = document.querySelector(".tags-container")
tagsContainer.append(CreateRect(tagsContainer,options))
