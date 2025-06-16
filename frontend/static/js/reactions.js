document.querySelectorAll('.like-button, .dislike-button, .like-comment-button, .dislike-comment-button').forEach((button) => {
  button.addEventListener('click', function (event) {
    event.preventDefault();

    const postId = button.getAttribute('data-posted-id');
    const reaction = button.getAttribute('data-reaction');

    fetch('/reaction', {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      credentials: 'include',
      body: `post_id=${postId}&reaction=${reaction}`,
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.success) {
          // Find the parent container for this post
          const postActions = button.closest('.post-actions') || button.closest('.comment-actions');
          if (!postActions) {
            console.error('Unable to find post-actions or comment-actions container.');
            return;
          }
          // Update the like and dislike counts dynamically
          const likeCountSpan = postActions.querySelector('.like-count');
          const dislikeCountSpan = postActions.querySelector('.dislike-count');

          likeCountSpan.textContent = data.likes;
          dislikeCountSpan.textContent = data.dislikes;
        } else {
          console.error('Failed to update reaction:', data.message);
        } 
      })
      .catch((err) => console.error(err));
  });
});