const essentialBtn = document.getElementById('essential');
const nonEssentialBtn = document.getElementById('nonEssential');
const output = document.getElementById('output');
const currentLoadSpan = document.getElementById('currentLoad');
const requestBtns = document.querySelectorAll('.request-btn');

let simulatedRequests = 0;
let simulatingLoad = false;

function updateLoadDisplay() {
  currentLoadSpan.textContent = simulatedRequests;
  currentLoadSpan.style.color = simulatedRequests > 50 ? '#e74c3c' : '#27ae60';
}

async function simulateLoad(requestCount) {
  if (simulatingLoad) return;

  simulatingLoad = true;
  const btn = event.target;
  const originalText = btn.textContent;

  requestBtns.forEach((b) => b.classList.remove('active'));
  btn.classList.add('active');
  btn.textContent = 'Simulating...';

  simulatedRequests = requestCount;
  updateLoadDisplay();

  const promises = [];
  for (let i = 0; i < requestCount; i++) {
    promises.push(
      fetch('/simulate-load', { method: 'POST' }).catch(() => {}) // Ignore errors
    );
  }

  btn.textContent = `Active (${requestCount} requests)`;

  // Wait for all requests to complete (they take 2 seconds each)
  try {
    await Promise.all(promises);
  } catch (e) {
    // Ignore errors
  }

  // Reset after simulation completes
  simulatedRequests = 0;
  updateLoadDisplay();
  btn.textContent = originalText;
  btn.classList.remove('active');
  simulatingLoad = false;
}

async function testEssentialService() {
  output.className = 'output';
  output.textContent = 'Testing essential service...';

  try {
    const res = await fetch('/essential');
    const text = await res.text();

    if (res.ok) {
      output.className = 'output success';
      output.textContent = `${text}`;
    } else {
      output.className = 'output error';
      output.textContent = `${text}`;
    }
  } catch (error) {
    output.className = 'output error';
    output.textContent = `${error.message}`;
  }
}

async function testNonEssentialService(event) {
  output.className = 'output';
  output.textContent = 'Testing non-essential service...';

  try {
    const res = await fetch('/non-essential');
    const text = await res.text();    

    if (res.status === 503 && simulatedRequests > 50) {
      output.className = 'output error';
      output.textContent = `Load shedding active - service temporarily unavailable due to high load.`;
    } else if (res.ok && simulatedRequests > 50) {
      output.className = 'output error';
      output.textContent = `Load shedding active - service temporarily unavailable due to high load. Priority given to essential services`;
    } else {
      output.className = 'output error';
      output.textContent = `(${res.status})\n${text}`;
    }
  } catch (error) {
    output.className = 'output error';
    output.textContent = `${error.message}`;
  }
}

requestBtns.forEach((btn) => {
  btn.addEventListener('click', (event) => {
    const count = parseInt(btn.dataset.count);
    simulateLoad(count);
  });
});

essentialBtn.addEventListener('click', testEssentialService);
nonEssentialBtn.addEventListener('click', testNonEssentialService);

updateLoadDisplay();
