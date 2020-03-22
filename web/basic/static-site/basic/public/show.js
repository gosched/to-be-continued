let url = 'https://images.unsplash.com/photo-1528035396761-8bd12d57efbd';

fetch(url)
    .then((response) => {
        return response.blob(); // Binary Large Object
    })
    .then((imgBlob) => {
    
    let img = document.createElement('IMG')
    img.src = URL.createObjectURL(imgBlob);
    
    let show = document.getElementById("show");
    show.appendChild(img);
  })