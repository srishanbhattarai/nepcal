class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/releases/download/v0.5.0/nepcal_0.5.0_darwin_amd64.tar.gz"
  version "0.5.0"
  sha256 "021e4bede0bf42d771b8cabaa3d55763c2fe180280baabdde2ffd5433b6e0235"

  def install
    bin.install "nepcal"
  end
end
