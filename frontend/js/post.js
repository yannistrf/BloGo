import { doLogout, getToken } from "./auth.js";
import { getPost } from "./api.js";

const urlParams = new URLSearchParams(window.location.search);
const postId = urlParams.get('post_id');

const token = getToken();
if (!token) {
  window.location.href = "/login.html";
}

document.getElementById("logoutBtn").addEventListener("click", () => {
    doLogout()
    window.location.replace("/login.html")
})

getPost(token, postId).then(post => {
    if (Object.keys(post).length === 0) {
      document.getElementById("post-title").textContent = "Invalid post id";
      return;
    }

    document.getElementById("post-title").textContent = post.title;
    document.getElementById("post-meta").textContent =
      `By ${post.author} • ${new Date(post.created_at).toLocaleString()}`;
    document.getElementById("post-content").textContent = post.content;
})