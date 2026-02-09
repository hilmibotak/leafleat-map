// Configuration using jscroot approach
export const config = {
    apiBaseUrl: 'http://localhost:8080/api',
    mapCenter: [20, 0],
    mapZoom: 2,
    updateInterval: 5000,
    packetUpdateInterval: 100
};

export const teams = [
    {
        id: 1,
        name: 'Phoenix Hackers',
        location: 'Tokyo, Japan',
        ip: '192.168.1.101',
        lat: 35.6762,
        lng: 139.6503,
        color: 'red',
        members: 5,
        score: 0,
        solved: 0,
        progress: 88
    },
    {
        id: 2,
        name: 'Cyber Ninjas',
        location: 'San Francisco, USA',
        ip: '192.168.1.102',
        lat: 37.7749,
        lng: -122.4194,
        color: 'orange',
        members: 5,
        score: 0,
        solved: 0,
        progress: 59
    },
    {
        id: 3,
        name: 'Binary Wolves',
        location: 'Berlin, Germany',
        ip: '192.168.1.103',
        lat: 52.5200,
        lng: 13.4050,
        color: 'green',
        members: 5,
        score: 0,
        solved: 0,
        progress: 88
    },
    {
        id: 4,
        name: 'Shadow Raiders',
        location: 'Singapore',
        ip: '192.168.1.104',
        lat: 1.3521,
        lng: 103.8198,
        color: 'yellow',
        members: 5,
        score: 0,
        solved: 0,
        progress: 59
    },
    {
        id: 5,
        name: 'Quantum Knights',
        location: 'Sydney, Australia',
        ip: '192.168.1.105',
        lat: -33.8688,
        lng: 151.2093,
        color: 'purple',
        members: 5,
        score: 0,
        solved: 0,
        progress: 69
    }
];
