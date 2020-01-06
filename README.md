# GoServerTemplate
Simple Go Server

Working default.nix 

{ lib, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
	name = "GoServerTemplate-${version}";
	version = "master-2020-01-06";

	modSha256 = "1pwmg48p1r670l83w5bwh0drc7mix8av4p2yhi1b5i5pw18i9i2g";

	src = fetchFromGitHub{
		owner = "rengglian";
		repo = "GoServerTemplate";
		rev = "e575bd82c11d99ac2aed6d7d55a60528277d9fc5";
		sha256 = "1vpm618l39pjh67k62h2ksxr30y7d98xcr03xgdddpyi0rbapq97";
	};

	subPackages = [ "." ];

	postInstall = ''
		cp -rf $src/addons $out/bin
		cp -rf $src/templates/ $out/bin
		cp -rf $src/config/ $out/bin
	'';

	meta = with lib; {
		description = "Simple Go Server Template.";
		homepage = https://github.com/rengglian/GoServerTemplate;
		license = licenses.mit;
		maintainers = with maintainers; [ rengglian ];
		platforms = platforms.linux ++ platforms.darwin;
	};
}
