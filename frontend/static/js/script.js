const essentialBtn = document.getElementById('essential');
const nonEssentialBtn = document.getElementById('nonEssential');
const output = document.getElementById('output');

async function essentialService() {
  const res = await fetch('/essential');
  const text = await res.text();
  output.textContent = text;
}

async function nonEssentialService() {
  const res = await fetch('/non-essential');
  const text = await res.text();

  if (res.status === 503) {
    output.textContent = 'Feature currently unavailable. Try again later';
  } else {
    output.textContent = text;
  }
}

essentialBtn.addEventListener('click', essentialService);
nonEssentialBtn.addEventListener('click', nonEssentialService);
