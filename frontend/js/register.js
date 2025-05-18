import { register } from "./api.js";

document.getElementById("registerForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;
  const password_confirm = document.getElementById("confirm-password").value;

  if (password !== password_confirm) {
    alert("Passwords dont match");
    return
  }

  const ok = await register(username, password);
  if (ok) {
    window.location.href = "login.html";
  } else {
    alert("Registration failed");
  }
});