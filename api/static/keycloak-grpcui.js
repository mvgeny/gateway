async function fillAuthorizationMetadata(){
    const keycloak = new Keycloak({
        url: 'http://127.0.0.1:8085',
        realm: 'gateway',
        clientId: 'myclient'
    });
    const authenticated = await keycloak.init({ onLoad: 'login-required', checkLoginIframe: 
false })
    if(authenticated){
        document.getElementsByClassName("name")[1].value = "Authorization"
        document.getElementsByClassName("value")[1].value = `Bearer ${keycloak.token}`
        setInterval(async () => {
            await keycloak.updateToken(60);
            document.getElementsByClassName("value")[1].value = `Bearer ${keycloak.token}`
        }, 55000);
    }
}
fillAuthorizationMetadata();
