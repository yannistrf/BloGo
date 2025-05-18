import { getPosts } from "./api.js";

getPosts().then(posts => {
    const container = document.getElementById('postsContainer');
    container.innerHTML = '';
    posts.reverse().forEach(post => {
        const postElement = document.createElement('div');
        // postElement.className = 'text-center';
        postElement.innerHTML = `
        <div class="card shadow-sm" style="max-width: 800px;min-width: 500px;">
            <div class="card-body">
                <h5 class="card-title">${post.title}</h5>
                <p class="text-muted position-absolute top-0 end-0 m-3">Written by: Johcscevevn<\p>
                <p class="card-text text-truncate">${post.content}</p>
            </div>
        </div>
        `;
        container.appendChild(postElement);
    });
})
