document.querySelectorAll('.post-time time').forEach((el) => {
  const rawTimestamp = el.getAttribute('datetime');
  const parsedTimestamp = new Date(rawTimestamp.replace(' +0000 UTC', 'Z'));

  if (!isNaN(parsedTimestamp)) {
    el.innerText = formatTimestamp(parsedTimestamp);
  } else {
    console.warn(`Invalid date format: ${rawTimestamp}`);
  }
});

function formatTimestamp(timestamp) {
  const now = new Date();
  const pastTime = new Date(timestamp);
  const timeDifference = Math.floor((now - pastTime) / 1000);

  const units = [
    { label: 'year', seconds: 31536000 },
    { label: 'month', seconds: 2592000 },
    { label: 'week', seconds: 604800 },
    { label: 'day', seconds: 86400 },
    { label: 'hour', seconds: 3600 },
    { label: 'minute', seconds: 60 },
    { label: 'second', seconds: 1 }
  ];

  for (const unit of units) {
    const interval = Math.floor(timeDifference / unit.seconds);
    if (interval >= 1) {
      return interval === 1 ? `1 ${unit.label} ago` : `${interval} ${unit.label}s ago`;
    }
  }

  return 'just now'; // In case the time difference is 0
}

fetch('/posts')
  .then((response) => response.json())
  .then((posts) => {
    posts.forEach((post) => {
      const timeElement = document.querySelector(
        `.post-time time[datetime="${post.created_on}"]`
      );

      if (timeElement) {
        const timestamp = new Date(post.created_on);
        timeElement.innerText = formatTimestamp(timestamp);
      } else {
        console.warn(
          `No matching <time> element found for timestamp: ${post.created_on}`
        );
      }
    });
  })
  .catch((error) => console.error('Error fetching posts:', error));
