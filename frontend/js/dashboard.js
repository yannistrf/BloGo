import { getPosts, getMyPosts, createPost, getQueryPosts } from "./api.js";
import { doLogout, getToken } from "./auth.js";

const token = getToken();
if (!token) {
  window.location.href = "/login.html";
}

getPosts(token).then(posts => {
    createPostElements(posts);
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
        createPostElements(posts);
    });
})

document.getElementById("queryForm").addEventListener("submit", async (e) => {
    e.preventDefault()
    const query = document.getElementById("query").value
    if (query === "") {
        return
    }
    document.getElementById("header").textContent = `Search results for: ${query}`

    const container = document.getElementById('postsContainer');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }

    getQueryPosts(token, query).then(posts => {
        createPostElements(posts);
    });
})

function createPostElements(posts) {
    const container = document.getElementById('postsContainer');
    posts.reverse().forEach(post => {
        const date = new Date(post.created_at)
        const dateString = date.toLocaleDateString();
        const timeString = date.toLocaleTimeString();

        const postElement = document.createElement('div');
        postElement.innerHTML = `
        <div class="card shadow-sm" style="max-width: 800px;min-width: 500px;">
            <div class="card-body">
                <h5 class="card-title">${post.title}</h5>
                <p class="text-muted position-absolute top-0 end-0 m-3">Written by: ${post.author}</p>
                <p class="card-text" style="white-space: pre-wrap;">${post.content}</p>
                <p class="text-muted mb-0">Created at: ${dateString} (${timeString})</p>
            </div>
        </div>
        `;
        container.appendChild(postElement);
    });
}

