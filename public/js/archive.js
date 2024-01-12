let monthTotalContainer = document.querySelectorAll(".monthlist-container")
monthTotalContainer.forEach(container=>{
   let totalPages =  container.querySelectorAll("li").length
   container.querySelector(".month-total").textContent = totalPages
})


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

function CreateColumnLine(container,options){
    const computedStyle = getComputedStyle(container)

    const height = computedStyle.height
    const width = computedStyle.fontSize
    const intheight = parseInt(height)
    const intwidth = parseInt(width)

    var svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("width", "1em");
    svg.setAttribute("height", "calc(100% - 1rem)");

    const rc = rough.svg(svg)
    let node = rc.line(intwidth/2,0,intwidth/2,intheight,options)
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

let monthContainers = document.querySelectorAll(".monthlist-container")
monthContainers.forEach(container=>{
    container.append(CreateColumnLine(container,options))
})