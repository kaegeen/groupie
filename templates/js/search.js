const searchBox = document.getElementById('search-input');
const creationDateS = document.getElementById('creationdatestart');
const creationDateE = document.getElementById('creationdateend');
const valueStart = document.getElementById('valueStart');
const valueEnd = document.getElementById('valueEnd');
const firstAlbumS = document.getElementById('firstalbumstart');
const firstAlbumE = document.getElementById('firstalbumend');
const valueAlbumStart = document.getElementById('valueAlbumStart');
const valueAlbumEnd = document.getElementById('valueAlbumEnd');

let timeoutId;

searchBox.addEventListener('keyup', debounce(search, 500));
creationDateS.addEventListener('input', handleDateChange);
creationDateE.addEventListener('input', handleDateChange);
firstAlbumS.addEventListener('input', handleAlbumChange);
firstAlbumE.addEventListener('input', handleAlbumChange);

function handleDateChange(e) {
    const start = parseInt(creationDateS.value, 10);
    const end = parseInt(creationDateE.value, 10);
    
    if (start > end) {
        creationDateE.value = start;
        valueEnd.innerText = start;
    } else {
        valueStart.innerText = start;
        valueEnd.innerText = end;
    }
    search();
}

function handleAlbumChange(e) {
    const start = parseInt(firstAlbumS.value, 10);
    const end = parseInt(firstAlbumE.value, 10);
    
    if (start > end) {
        firstAlbumE.value = start;
        valueAlbumEnd.innerText = start;
    } else {
        valueAlbumStart.innerText = start;
        valueAlbumEnd.innerText = end;
    }
    search();
}

function debounce(func, delay) {
    return function (...args) {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => func(...args), delay);
    };
}

function search() {
    const keywords = searchBox.value.toLowerCase().trim().split('-')[0].trim();
    const creationDateStart = parseInt(creationDateS.value, 10);
    const creationDateEnd = parseInt(creationDateE.value, 10);
    const firstAlbumStart = parseInt(firstAlbumS.value, 10);
    const firstAlbumEnd = parseInt(firstAlbumE.value, 10);

    const artistName = document.getElementById('search-artist');
    const artistList = document.getElementsByClassName('opt-artist');

    Array.from(artistList).forEach(elm => elm.remove());

    Array.from(document.getElementsByClassName('main-article')).forEach(elm => {
        const allNames = elm.getAttribute('data-allNames').toLowerCase();
        const authorMemb = elm.getAttribute('data-authorMemb');
        const creationDate = parseInt(elm.getAttribute('data-creationDate'), 10);
        const firstAlbumDate = parseInt(elm.getAttribute('data-firstAlbumDate').slice(-4), 10);

        if (allNames.includes(keywords) &&
            creationDate >= creationDateStart &&
            creationDate <= creationDateEnd &&
            firstAlbumDate >= firstAlbumStart &&
            firstAlbumDate <= firstAlbumEnd) {
            
            elm.style.display = "block";
            const names = Array.from(new Set(allNames.split("|").filter(name => name.includes(keywords))));
            
            names.forEach(name => {
                const opt = document.createElement('option');
                opt.value = capitalize(`${name} - ${authorMemb}`);
                opt.className = "opt-artist";
                artistName.appendChild(opt);
            });
        } else {
            elm.style.display = "none";
        }
    });
}

function capitalize(str) {
    return str.replace(/(^\w{1})|(\s+\w{1})/g, letter => letter.toUpperCase());
}
