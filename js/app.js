import { initMap } from './map.js';
import { renderProgressBars, renderTeamCards } from './teams.js';
import { config } from './config.js';

// Initialize clock
function updateClock() {
    const now = new Date();
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');
    
    const clockElement = document.getElementById('clock');
    if (clockElement) {
        clockElement.textContent = `${hours}:${minutes}:${seconds}`;
    }
}

// Update packet counter
function updatePacketCounter() {
    const packetElement = document.getElementById('packetsPerSec');
    if (packetElement) {
        const packets = Math.floor(Math.random() * 50) + 100; // Random between 100-150
        packetElement.textContent = packets;
    }
}

// Update active nodes count
function updateActiveNodes() {
    const nodesElement = document.getElementById('activeNodesCount');
    if (nodesElement) {
        nodesElement.textContent = '5';
    }
}

// Simulate real-time updates
function simulateRealTimeUpdates() {
    setInterval(() => {
        updatePacketCounter();
        // Add more real-time updates here if needed
    }, config.packetUpdateInterval);
}

// Initialize application
function init() {
    console.log('🚀 Cyber Map CTF Initializing...');
    
    // Initialize map
    initMap();
    
    // Render UI elements
    renderProgressBars();
    renderTeamCards();
    
    // Start clock
    updateClock();
    setInterval(updateClock, 1000);
    
    // Update active nodes
    updateActiveNodes();
    
    // Start real-time updates
    simulateRealTimeUpdates();
    
    // Initialize sidebar toggle
    initSidebarToggle();
    
    // Initialize navigation
    initNavigation();
    
    console.log('✅ Cyber Map CTF Ready!');
}

// Sidebar toggle for mobile
function initSidebarToggle() {
    const menuToggle = document.getElementById('menuToggle');
    const sidebar = document.querySelector('.sidebar');
    const mainContent = document.querySelector('.main-content');
    
    if (menuToggle && sidebar) {
        menuToggle.addEventListener('click', function() {
            sidebar.classList.toggle('open');
            menuToggle.classList.toggle('active');
        });
        
        // Close sidebar when clicking outside
        document.addEventListener('click', function(e) {
            if (window.innerWidth <= 968) {
                if (!sidebar.contains(e.target) && !menuToggle.contains(e.target)) {
                    sidebar.classList.remove('open');
                    menuToggle.classList.remove('active');
                }
            }
        });
        
        // Close sidebar when clicking nav items on mobile
        const navItems = sidebar.querySelectorAll('.nav-item');
        navItems.forEach(item => {
            item.addEventListener('click', function() {
                if (window.innerWidth <= 968) {
                    sidebar.classList.remove('open');
                    menuToggle.classList.remove('active');
                }
            });
        });
    }
}

// Navigation between sections
function initNavigation() {
    const navItems = document.querySelectorAll('.nav-item');
    const sections = {
        'map': document.getElementById('section-map'),
        'teams': document.getElementById('section-teams'),
        'stats': document.getElementById('section-stats')
    };

    navItems.forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            
            // Remove active class from all nav items
            navItems.forEach(nav => nav.classList.remove('active'));
            
            // Add active class to clicked item
            item.classList.add('active');
            
            // Get the page to show
            const page = item.getAttribute('data-page');
            
            // Hide all sections
            Object.values(sections).forEach(section => {
                if (section) section.style.display = 'none';
            });
            
            // Show selected section
            if (sections[page]) {
                sections[page].style.display = 'block';
                
                // If switching to map, invalidate size to fix display issues
                if (page === 'map' && window.map) {
                    setTimeout(() => {
                        window.map.invalidateSize();
                    }, 100);
                }
            }
            
            // Close sidebar on mobile after selection
            if (window.innerWidth <= 968) {
                document.querySelector('.sidebar').classList.remove('open');
                const menuToggle = document.getElementById('menuToggle');
                if (menuToggle) menuToggle.classList.remove('active');
            }
        });
    });
}

// Wait for DOM to load
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
} else {
    init();
}
