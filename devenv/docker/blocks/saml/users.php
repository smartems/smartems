<?php
$config = array(
    'admin' => array(
        'core:AdminPassword',
    ),
    'example-userpass' => array(
        'exampleauth:UserPass',
        'saml-admin:grafana' => array(
            'groups' => array('admins'),
            'email' => 'saml-admin@smartEvo.de',
        ),
        'saml-editor:grafana' => array(
            'groups' => array('editors'),
            'email' => 'saml-editor@smartEvo.de',
        ),
        'saml-viewer:grafana' => array(
            'groups' => array(),
            'email' => 'saml-viewer@smartEvo.de',
        ),
    ),
);
