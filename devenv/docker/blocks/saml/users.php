<?php
$config = array(
    'admin' => array(
        'core:AdminPassword',
    ),
    'example-userpass' => array(
        'exampleauth:UserPass',
        'saml-admin:smartems' => array(
            'groups' => array('admins'),
            'email' => 'saml-admin@smartEvo.de',
        ),
        'saml-editor:smartems' => array(
            'groups' => array('editors'),
            'email' => 'saml-editor@smartEvo.de',
        ),
        'saml-viewer:smartems' => array(
            'groups' => array(),
            'email' => 'saml-viewer@smartEvo.de',
        ),
    ),
);
