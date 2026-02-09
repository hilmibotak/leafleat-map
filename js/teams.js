import { teams } from './config.js';

export function renderProgressBars() {
    const container = document.getElementById('progressBars');
    if (!container) return;

    container.innerHTML = teams.map(team => `
        <div class="progress-item">
            <div class="progress-header">
                <div class="progress-team-info">
                    <div class="progress-indicator" style="background-color: ${getColorHex(team.color)};"></div>
                    <span class="progress-team-name">${team.name}</span>
                </div>
                <span class="progress-percentage-value">${team.progress}%</span>
            </div>
            <div class="progress-bar-wrapper">
                <div class="progress-bar-fill" style="width: ${team.progress}%; background: linear-gradient(90deg, ${getColorHex(team.color)}, #b537f2);"></div>
            </div>
        </div>
    `).join('');
}

export function renderTeamCards() {
    const container = document.getElementById('teamsContainer');
    if (!container) return;

    container.innerHTML = teams.map(team => `
        <div class="team-card ${team.color}">
            <div class="team-card-header">
                <div class="team-title">${team.name}</div>
                <div class="team-badge">TEAM ${team.id}</div>
            </div>
            
            <div class="team-info">
                <div class="team-detail">
                    <svg viewBox="0 0 24 24" fill="currentColor">
                        <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7zm0 9.5c-1.38 0-2.5-1.12-2.5-2.5s1.12-2.5 2.5-2.5 2.5 1.12 2.5 2.5-1.12 2.5-2.5 2.5z"/>
                    </svg>
                    ${team.location}
                </div>
                <div class="team-detail">
                    <svg viewBox="0 0 24 24" fill="currentColor">
                        <path d="M17 7h-4v1.9h4c1.71 0 3.1 1.39 3.1 3.1 0 1.43-.98 2.63-2.31 2.98l.02.02H14v1.9h3.81c2.39-.35 4.19-2.4 4.19-4.9C22 9.39 19.61 7 17 7z"/>
                        <path d="M11 7H7c-2.76 0-5 2.24-5 5s2.24 5 5 5h4v-1.9H7c-1.71 0-3.1-1.39-3.1-3.1 0-1.71 1.39-3.1 3.1-3.1h4V7z"/>
                        <path d="M8 11h8v2H8z"/>
                    </svg>
                    ${team.ip}
                </div>
            </div>
            
            <div class="team-stats-inline">
                <div class="stat-inline">
                    <span class="stat-inline-value">${team.members}</span>
                    <span class="stat-inline-label">Members</span>
                </div>
                <div class="stat-inline">
                    <span class="stat-inline-value">${team.score}</span>
                    <span class="stat-inline-label">Score</span>
                </div>
                <div class="stat-inline">
                    <span class="stat-inline-value">${team.solved}</span>
                    <span class="stat-inline-label">Solved</span>
                </div>
            </div>
        </div>
    `).join('');
}

export function updateTeamScore(teamId, score, solved) {
    const team = teams.find(t => t.id === teamId);
    if (team) {
        team.score = score;
        team.solved = solved;
        renderTeamCards();
    }
}

export function updateTeamProgress(teamId, progress) {
    const team = teams.find(t => t.id === teamId);
    if (team) {
        team.progress = progress;
        renderProgressBars();
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

export function getTeams() {
    return teams;
}
