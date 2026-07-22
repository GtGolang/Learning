//replace the label content by button's content
function changeLabel(buttonElement){
    //find the label element by its ID 
    const label=document.getElementById("myLabel");
    //change the label content for this button content  
    label.textContent=buttonElement.textContent;
}