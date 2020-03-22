document.addEventListener('submit', e => {
    e.preventDefault();
    const form = e.target;
    
    fetch(form.action, {
        method: form.method,
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'multipart/form-data'
        },
    body: new FormData(form);
    }).then(function checkStatus(response) {
        if (response.status >= 200 && response.status < 300) {
            return response.json();
        } else {
            var error = new Error(response.statusText)
            error.response = response;
            throw error;
        }
       })
       .then(function(data) {

        }).catch(function(error) {
           console.log('request failed', error);
           return error.response.json();
       }).then(function(errorData){

       });
});


