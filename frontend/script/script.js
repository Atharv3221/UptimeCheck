const submit = document.querySelector("button");
const div = document.getElementsByTagName("div")[0];

submit.addEventListener("click", () => {
  const inputUrl = document.getElementsByTagName("input")[0];
  const message = document.createElement("p");
  message.textContent = "Submitted successfully";
  message.style.color = "green";
  if(!inputUrl.checkValidity()) {
    message.textContent = "URL is not valid"
    message.style.color = "red";
    div.appendChild(message);
  } else {
    const url = inputUrl.value;
    console.log(url);
    div.appendChild(message);
  }
  setTimeout(() => {
    message.remove(); 
  }, 3000);
});
