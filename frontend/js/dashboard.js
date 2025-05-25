import { getPosts, getMyPosts, createPost } from "./api.js";
import { doLogout, getToken } from "./auth.js";

const token = getToken();
if (!token) {
  window.location.href = "/login.html";
}

getPosts(token).then(posts => {
    const container = document.getElementById('postsContainer');
    posts.reverse().forEach(post => {
        const postElement = document.createElement('div');
        postElement.innerHTML = `
        <div class="card shadow-sm" style="max-width: 800px;min-width: 500px;">
            <div class="card-body">
                <h5 class="card-title">${post.title}</h5>
                <p class="text-muted position-absolute top-0 end-0 m-3">Written by: Johcscevevn<\p>
                <p class="card-text" style="white-space: pre-wrap;">${post.content}</p>
            </div>
        </div>
        `;
        container.appendChild(postElement);
    });
})

document.getElementById("newPostForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const postTitle = document.getElementById("postTitle").value;
    const postContent = document.getElementById("postContent").value;

    const ok = await createPost(token, postTitle, postContent);
    if (ok) {
        document.getElementById("postTitle").value = "";
        document.getElementById("postContent").value = "";
        location.reload();
    } else {
        alert("New post creation failed");
    }
})

document.getElementById("logoutBtn").addEventListener("click", () => {
    doLogout()
    window.location.replace("/login.html")
})

document.getElementById("myPostsBtn").addEventListener("click", () => {
    
    document.getElementById("header").textContent = "Your posts"

    const container = document.getElementById('postsContainer');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }

    getMyPosts(token).then(posts => {
        const container = document.getElementById('postsContainer');
        posts.reverse().forEach(post => {
            const postElement = document.createElement('div');
            postElement.innerHTML = `
            <div class="card shadow-sm" style="max-width: 800px;min-width: 500px;">
                <div class="card-body">
                    <h5 class="card-title">${post.title}</h5>
                    <p class="text-muted position-absolute top-0 end-0 m-3">Written by: Johcscevevn<\p>
                    <p class="card-text" style="white-space: pre-wrap;">${post.content}</p>
                </div>
            </div>
            `;
            container.appendChild(postElement);
        });
    })
})

