const inpFile = document.getElementById("inpFile");
const previewImage = document.getElementById("profileImage");

inpFile.addEventListener("change", function(){
    const file = this.files[0];
    if(file){
        const reader = new FileReader();
        previewImage.style.display = "block";
        reader.addEventListener("load", function() {
            previewImage.setAttribute("src", this.result);
        })
        if(file.size < 20 *1024*1024){
            reader.readAsDataURL(file);
        }else{
            window.alert("The file is too big");
            inpFile.value = ""
        }
    }else{
        previewImage.setAttribute("src","/static/images/default-avatar.jpg");
    }
})
$('#newAva').submit(function(){
    if(inpFile.value) return true
    return false
})