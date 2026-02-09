import { config } from './config.js';

// API functions using fetch (vanilla JS)
export async function fetchTeams() {
    try {
        const response = await fetch(`${config.apiBaseUrl}/teams`);
        if (!response.ok) throw new Error('Failed to fetch teams');
        return await response.json();
    } catch (error) {
        console.error('Error fetching teams:', error);
        return null;
    }
}

export async function fetchTeamById(teamId) {
    try {
        const response = await fetch(`${config.apiBaseUrl}/teams/${teamId}`);
        if (!response.ok) throw new Error('Failed to fetch team');
        return await response.json();
    } catch (error) {
        console.error('Error fetching team:', error);
        return null;
    }
}

export async function updateTeamStatus(teamId, data) {
    try {
        const response = await fetch(`${config.apiBaseUrl}/teams/${teamId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) throw new Error('Failed to update team');
        return await response.json();
    } catch (error) {
        console.error('Error updating team:', error);
        return null;
    }
}

export async function fetchChallenges() {
    try {
        const response = await fetch(`${config.apiBaseUrl}/challenges`);
        if (!response.ok) throw new Error('Failed to fetch challenges');
        return await response.json();
    } catch (error) {
        console.error('Error fetching challenges:', error);
        return null;
    }
}

export async function submitFlag(teamId, challengeId, flag) {
    try {
        const response = await fetch(`${config.apiBaseUrl}/submit`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ teamId, challengeId, flag })
        });
        if (!response.ok) throw new Error('Failed to submit flag');
        return await response.json();
    } catch (error) {
        console.error('Error submitting flag:', error);
        return null;
    }
}

export async function fetchNetworkActivity() {
    try {
        const response = await fetch(`${config.apiBaseUrl}/activity`);
        if (!response.ok) throw new Error('Failed to fetch network activity');
        return await response.json();
    } catch (error) {
        console.error('Error fetching network activity:', error);
        return null;
    }
}
