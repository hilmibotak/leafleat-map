import { config, teams } from './config.js';

let map;
let markers = {};
let polylines = [];

export function initMap() {
    // Initialize Leaflet map
    map = L.map('map', {
        center: config.mapCenter,
        zoom: config.mapZoom,
        zoomControl: true,
        attributionControl: false
    });

    // Use dark blue/purple tile layer for cyber theme
    L.tileLayer('https://tiles.stadiamaps.com/tiles/alidade_smooth_dark/{z}/{x}/{y}{r}.png', {
        maxZoom: 19
    }).addTo(map);

    // Add team markers
    teams.forEach(team => {
        addTeamMarker(team);
    });

    // Draw connections between teams
    drawConnections();

    // Update coordinates on mouse move
    map.on('mousemove', function(e) {
        updateCoordinates(e.latlng.lat, e.latlng.lng);
    });

    // Animate connections
    animateConnections();

    return map;
}

function addTeamMarker(team) {
    const markerIcon = L.divIcon({
        className: 'custom-marker',
        html: `<div class="team-marker" style="color: ${getColorHex(team.color)}; border-color: ${getColorHex(team.color)};">
                ${team.id}
              </div>`,
        iconSize: [40, 40],
        iconAnchor: [20, 20]
    });

    const marker = L.marker([team.lat, team.lng], { icon: markerIcon })
        .addTo(map)
        .bindPopup(`
            <div style="color: #000; font-weight: 600;">
                <strong>${team.name}</strong><br>
                ${team.location}<br>
                IP: ${team.ip}
            </div>
        `);

    markers[team.id] = marker;
}

function drawConnections() {
    // Draw lines between teams to simulate network connections
    const teamPairs = [
        [teams[0], teams[1]],
        [teams[1], teams[2]],
        [teams[2], teams[3]],
        [teams[3], teams[4]],
        [teams[0], teams[3]]
    ];

    teamPairs.forEach(([team1, team2]) => {
        const polyline = L.polyline(
            [[team1.lat, team1.lng], [team2.lat, team2.lng]],
            {
                color: '#4da6ff',
                weight: 2,
                opacity: 0.5,
                dashArray: '5, 10'
            }
        ).addTo(map);

        polylines.push(polyline);
    });
}

function animateConnections() {
    let offset = 0;
    setInterval(() => {
        offset = (offset + 1) % 15;
        polylines.forEach(polyline => {
            polyline.setStyle({
                dashOffset: offset
            });
        });
    }, 100);
}

function updateCoordinates(lat, lng) {
    const coordElement = document.getElementById('mapCoordinates');
    if (coordElement) {
        coordElement.textContent = `LAT: ${lat.toFixed(4)} | LNG: ${lng.toFixed(4)}`;
    }
}

function getColorHex(colorName) {
    const colors = {
        red: '#ff4444',
        orange: '#ff9944',
        green: '#00ff88',
        yellow: '#ffdd44',
        purple: '#aa44ff'
    };
    return colors[colorName] || '#4da6ff';
}

export function updateMarkerPosition(teamId, lat, lng) {
    if (markers[teamId]) {
        markers[teamId].setLatLng([lat, lng]);
    }
}

export function getMap() {
    return map;
}
