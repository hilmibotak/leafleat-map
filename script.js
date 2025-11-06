// Membuat peta dengan lokasi pusat di Alun-Alun Bandung
// Format: L.map('id-element').setView([latitude, longitude], zoom-level)
var map = L.map('map').setView([-6.9219, 107.6070], 15); // Pusatkan di Alun-Alun Bandung

// Menambahkan tilelayer dari OpenStreetMap (gratis dan open-source)
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
  maxZoom: 19, // Angka maksimal zoom (19 sudah cukup)
  attribution: '© OpenStreetMap' // Hak cipta ke penyedia peta
}).addTo(map); // Tambahkan layer peta ke objek map

// Menambahkan marker pada Alun-Alun Bandung
var marker = L.marker([-6.9219, 107.6070]).addTo(map);

// Tambahkan popup (balon info) pada marker
marker.bindPopup("<b>Alun-Alun Bandung</b><br>Tempat ikonik di pusat kota Bandung.").openPopup();
