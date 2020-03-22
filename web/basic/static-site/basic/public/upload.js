document.addEventListener('submit', e => {
  e.preventDefault();
  const form = e.target;

  fetch(form.action, {
    method: form.method,
    headers: {
       'Accept': 'application/json',
       'Content-Type': 'multipart/form-data'
    },
    body: new FormData(form),
  })

});


// application/x-www-form-urlencoded
// multipart/form-data
// text/plain
