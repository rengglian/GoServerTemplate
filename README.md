# GoServerTemplate
Simple Go Server


{ lib, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
  name = "GoServerTemplate-${version}";
  version = "1.0.0";

  modSha256 = "1pwmg48p1r670l83w5bwh0drc7mix8av4p2yhi1b5i5pw18i9i2g";

  src = fetchFromGitHub{
    owner = "rengglian";
    repo = "GoServerTemplate";
    rev = "v${version}";
    sha256 = "1vpm618l39pjh67k62h2ksxr30y7d98xcr03xgdddpyi0rbapq97";
  };

  subPackages = [ "." ];

  meta = with lib; {
    description = "Simple Go Server Template.";
    homepage = https://github.com/rengglian/GoServerTemplate;
    license = licenses.mit;
    maintainers = with maintainers; [ rengglian ];
    platforms = platforms.linux ++ platforms.darwin;
  };
}
