import { login } from "./api.js";
import { saveToken } from "./auth.js"

document.getElementById("loginForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;
  const res = await login(username, password);
  if (res.token) {
    saveToken(res.token);
    window.location.href = "dashboard.html";
  } else {
    alert("Login failed");
  }
});