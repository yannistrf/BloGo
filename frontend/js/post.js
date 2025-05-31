import { doLogout, getToken } from "./auth.js";
import { getPost, publishComment } from "./api.js";

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

    createCommentElements(post);
})

document.getElementById('commentForm').addEventListener('submit', async (e) => {
  e.preventDefault()
  const commentContent = document.getElementById('commentContent').value;
  const ok = await publishComment(token, postId, commentContent);
  if (ok) {
    document.getElementById("commentContent").value = "";
    location.reload();
  } else {
    alert("Couldn't publish comment");
  }
})

function createCommentElements(post) {
  const container = document.getElementById('commentsContainer');

  post.comments.forEach(comment => {
    const commentElement = document.createElement('div');
    const date = new Date(comment.created_at)
    const dateString = date.toLocaleDateString();
    const timeString = date.toLocaleTimeString();
    commentElement.innerHTML = `
    <div class="card shadow-sm my-2">
      <div class="card-body">
      <div class="d-flex justify-content-between flex-wrap">
        <p class="text-muted">By: ${comment.author}</p>
        <p class="text-muted mb-0">${dateString} (${timeString})</p>
      </div>  
        <p class="card-text" style="white-space: pre-wrap;">${comment.content}</p>
      </div>
    </div>
    `;
    container.appendChild(commentElement);
  });
}
